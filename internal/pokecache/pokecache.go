package pokecache

import (
	"sync"
	"time"
)
type Cache struct {
	Cache map[string]cacheEntry
	sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) Set(key string, value []byte) {
	c.Lock()
	defer c.Unlock()

	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()

	entry, ok := c.Cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.Lock()
		now := time.Now()
		for key, val := range c.Cache {
			if now.Sub(val.createdAt) > interval {
				delete(c.Cache, key)
			}
		}
		c.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Cache: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return cache
}