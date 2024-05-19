package repl

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/shahank42/poketext/internal/app"
	"github.com/shahank42/poketext/internal/pokeapi"
)

func cmd_Help() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range CommandMap {
		fmt.Printf("\t%s:\t%s\n", command.name, command.description)
	}

	return nil
}

func cmd_Exit() error {
	os.Exit(0)
	return nil
}

func cmd_Map() error {
	if app.Config.NextLocationUrl == "" {
		return errors.New("No more locations! Try going back using `mapb`")
	}

	map_data, err := pokeapi.GetMap(app.Config.NextLocationUrl)
	if err != nil {
		return err
	}

	for _, location := range map_data.Results {
		fmt.Println(fmt.Sprintf("-* %s", location.Name))
	}

	app.Config.NextLocationUrl = map_data.Next
	app.Config.PreviousLocationUrl = map_data.Previous

	return nil
}

func cmd_MapBack() error {
	if app.Config.PreviousLocationUrl == "" {
		return errors.New("No more locations! Try going ahead using `map`")
	}

	map_data, err := pokeapi.GetMap(app.Config.PreviousLocationUrl)
	if err != nil {
		return err
	}

	for _, location := range map_data.Results {
		fmt.Println(fmt.Sprintf("-* %s", location.Name))
	}

	app.Config.NextLocationUrl = map_data.Next
	app.Config.PreviousLocationUrl = map_data.Previous

	return nil
}

func cmd_Explore(location string) error {
	exploreData, err := pokeapi.GetExplore(location)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Exploring %s...\nFound Pokemon:", location))
	for _, pokemon := range exploreData.PokemonEncounters {
		fmt.Println(fmt.Sprintf(" - %s", pokemon.Pokemon.Name))
	}

	return nil
}

func cmd_Catch(pokemon string) error {
	pokemonSpeciesData, err := pokeapi.GetPokemonSpeciesData(pokemon)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Throwing a pokeball at %s...", pokemonSpeciesData.Name))
	isCaught := rand.Intn(256) <= pokemonSpeciesData.CaptureRate
	if !isCaught {
		fmt.Println(fmt.Sprintf("%s escaped!", pokemonSpeciesData.Name))
		return nil
	}

	fmt.Println(fmt.Sprintf("%s was caught! Adding to Pokedex...", pokemonSpeciesData.Name))
	pokemonData, err := pokeapi.GetPokemonData(pokemon)
	if err != nil {
		return err
	}

	app.Config.Pokedex[pokemonData.Name] =
		app.PokedexEntry{
			Name:   pokemonData.Name,
			Height: pokemonData.Height,
			Weight: pokemonData.Weight,
			Stats:  pokemonData.Stats,
			Types:  pokemonData.Types,
		}

	return nil
}

func cmd_Inspect(pokemon string) error {
	pokemonData, ok := app.Config.Pokedex[pokemon]
	if !ok {
		fmt.Println(fmt.Sprintf("You haven't caught that Pokemon yet..."))
		return nil
	}

	fmt.Println(fmt.Sprintf("Name: %s", pokemonData.Name))
	fmt.Println(fmt.Sprintf("Height: %d", pokemonData.Height))
	fmt.Println(fmt.Sprintf("Weight: %d", pokemonData.Weight))
	fmt.Println("Base Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Println(fmt.Sprintf("\t- %s: %d", stat.Stat.Name, stat.BaseStat))
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonData.Types {
		fmt.Println(fmt.Sprintf("\t- %s", pokeType.Type.Name))
	}

	return nil
}

func cmd_Pokedex() error {
	fmt.Println(fmt.Sprintf("Your Pokedex:"))
	for _, pokemon := range app.Config.Pokedex {
		fmt.Println(fmt.Sprintf("\t- %s", pokemon.Name))
	}

	return nil
}
