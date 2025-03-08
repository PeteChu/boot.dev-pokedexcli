// Package pokecache provides a simple in-memory caching mechanism for storing
// Pokemon API response data with automatic expiration of entries.
//
// Usage:
//
//	// Create a new cache with a 5-minute expiration interval
//	cache := pokecache.NewCache(5 * time.Minute)
//
//	// Add data to the cache
//	cache.Add("pokemon/pikachu", responseData)
//
//	// Retrieve data from the cache
//	if data, found := cache.Get("pokemon/pikachu"); found {
//	    // Use the cached data
//	}
//
//	// Stop the cache cleaning goroutine when done
//	defer cache.Stop()
package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache represents an in-memory cache with automatic expiration of entries.
// It is safe for concurrent use by multiple goroutines.
type Cache struct {
	mu       sync.RWMutex
	cache    map[string]cacheEntry
	stopChan chan bool
}

// NewCache creates and initializes a new Cache instance with the specified expiration interval.
// It starts a background goroutine that periodically removes expired entries.
//
// The interval parameter specifies how long cache entries should live before being removed.
//
// When done with the cache, call the Stop method to clean up resources.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache:    make(map[string]cacheEntry),
		stopChan: make(chan bool),
	}
	go cache.reapLoop(interval)
	return cache
}

// Add stores a value in the cache with the specified key.
// If the key already exists, its value is overwritten.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves a value from the cache by its key.
// It returns the value and a boolean indicating whether the key was found.
// If the key is not in the cache, it returns nil and false.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cache, exist := c.cache[key]
	if !exist {
		return nil, false
	}

	return cache.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.cache {
				if t.Sub(entry.createdAt) > interval {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		case <-c.stopChan:
			return
		}
	}
}

// Stop terminates the background goroutine that cleans up expired cache entries.
// It's safe to call Stop multiple times.
// The cache should not be used after calling Stop.
func (c *Cache) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	select {
	case <-c.stopChan:
	default:
		close(c.stopChan)
	}
}
