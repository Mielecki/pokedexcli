package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
		mu: &sync.Mutex{},
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	newEntry := cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	cache.entries[key] = newEntry
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string)([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	res, ok := cache.entries[key]
	if ok {
		return res.val, true
	}
	var zero []byte
	return zero, false
}

func (cache *Cache) reapLoop() {
	for {
		cache.mu.Lock()
		for k, v := range cache.entries {
			if time.Since(v.createdAt).Seconds() > cache.interval.Seconds() {
				delete(cache.entries, k)
			}
		}
		cache.mu.Unlock()
		time.Sleep(cache.interval)
	}
}