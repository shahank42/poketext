package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cacheTable map[string]cacheEntry

type Cache struct {
	table    cacheTable
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	table := make(cacheTable)
	newCache := &Cache{table: table, interval: interval}
	go newCache.reapLoop()
	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.table[key] = cacheEntry{
		time.Now(),
		val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.table[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(time.Second) // Reap every second
	for range ticker.C {
		for url, entry := range cache.table {
			if time.Now().After(entry.createdAt.Add(cache.interval)) {
				delete(cache.table, url)
			}
		}
	}
}
