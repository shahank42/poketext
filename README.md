# PokeText

A Pokemon catching and (sorta) exploring CLI application.

Built for the corresponding <a href="https://www.boot.dev/learn/build-pokedex-cli">Boot.dev course</a> in an effort to learn the ins and outs of Golang by me.

# Installation
### Step 0: Getting PokeText Installed

Download an appropriate executable file from the <a href="https://github.com/shahank42/poketext/releases/latest">Releases</a> page. That's it.

### Step 1: Execute the executable from a CLI

Kinda self-explanatory, ain't it?

### Step 3: Gotta catch 'em all!

You're in the PokeText CLI. Remember to use the `help` command if you're ever stuck

# Compiling
Easy peasy.

```bash
  git clone https://github.com/shahank42/poketext.git
  cd poketext
  go build && ./poketext
```

This is for Loonix.
Do the corresponding thing if you're on Windows or some other OS.

# Why Even Bother With Something Like This?
That's exactly what I asked myself before I started. Like everyone else, I dismissed the idea of creating a CLI Pokedex because, well, it seemed too easy and not "cool enough". But then I realized that despite how much I keep touting about Golang I didn't have any proper project written in it.
Welp time to change that. I was gonna finish this seemingly easy project in one night.

It was more involved than I expected it to be. Not anything too difficult, but yeah definitely a good exercise of the brain.

Here are the things I learned along the way:
- Golang's package system for organizing code (which I absolutely love)
- Writing my own <a href="https://github.com/shahank42/poketext/blob/main/internal/pokecache/cache.go">cache</a> and <a href="https://github.com/shahank42/poketext/blob/main/internal/pokeapi/fetch.go">cache middleware</a> (yep, from scratch)
- How to think in the Golang way (can't really elaborate much on it, Go code has a certain feel to it which you can sense once you actually write a lot of it)
- Write code faster (I was practising using NeoVim so lmao)

Overall it's nice, maybe I can keep on developing it and add a bunch of cool features if I get enough willpower and motivation.

# Contributing
Lol idk why anyone would, but sure any contributions are welcome.

Maybe you're also learning Golang and are diving into open-source contributions? Well then PokeText can act as a cool first repository, just imagine contributing to some stranger's project :-)
