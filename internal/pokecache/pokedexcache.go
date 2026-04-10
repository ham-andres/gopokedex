package pokecache

import (

	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time 
	val []byte
}

type Cache struct {
	cachemap map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {

	currentcache := &Cache{
		cachemap: make(map[string]cacheEntry),
	}
	go currentcache.reapLoop(interval)
	return currentcache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:			 value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry, exists := c.cachemap[key]
	if exists {
		return entry.val, true 
	}
	return nil, false
}

// incomplete function
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	for range ticker.C {
		c.mu.Lock()
		
		cutoff := time.Now().Add(-interval)
		for key,val := range c.cachemap {
			if val.createdAt.Before(cutoff) {
				delete(c.cachemap,key)
		
			}
		}
		c.mu.Unlock()

	}
} 

