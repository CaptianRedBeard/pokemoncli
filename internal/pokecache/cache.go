package pokecache

import (
    "sync"
    "time"
)

type cacheEntry struct {
    createdAt time.Time 
    val       []byte    
}

type Cache struct {
    mu           sync.RWMutex  
    entries      map[string]cacheEntry
    reapInterval time.Duration 
}

func NewCache(reapInterval time.Duration) *Cache {
    cache := &Cache{
        entries:      make(map[string]cacheEntry),
        reapInterval: reapInterval,
    }
    go cache.reapLoop() 
    return cache
}

func (c *Cache) Add(key string, value []byte) {
    c.mu.Lock() 
    defer c.mu.Unlock()
    c.entries[key] = cacheEntry{
        createdAt: time.Now(),
        val:       value,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.RLock() 
    defer c.mu.RUnlock()
    entry, exists := c.entries[key]
    if !exists {
        return nil, false
    }
    return entry.val, true
}


func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.reapInterval)
    defer ticker.Stop()

    for {
        <- ticker.C
        c.mu.Lock()
        now := time.Now()

        for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.reapInterval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
