package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	Mu      *sync.RWMutex
}

type cacheFuncs interface {
	Add(key string, val []byte)
	Get(key string) (cache []byte, ok bool)
	reapLoop(interval time.Duration)
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(timeout time.Duration) Cache {
	c := Cache{
		Entries: make(map[string]cacheEntry),
		Mu:      &sync.RWMutex{},
	}
	go c.reapLoop(timeout)
	return c

}

func (c *Cache) Add(key string, val []byte) {

	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (cache []byte, ok bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	entry, ok := c.Entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {

		c.Mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()

	}
}
