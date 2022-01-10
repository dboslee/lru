package lru

import (
	"sync"

	"github.com/dboslee/lru/internal"
)

var _ internal.LRU[int, int] = &SyncCache[int, int]{}

// SyncCache is a threadsafe lru cache.
type SyncCache[K comparable, V any] struct {
	cache *Cache[K, V]
	mu    sync.RWMutex
}

// New initializes a new lru cache with the given capacity.
func NewSync[K comparable, V any](options ...CacheOption) *SyncCache[K, V] {
	return &SyncCache[K, V]{
		cache: New[K, V](options...),
	}
}

// Len is the number of key value pairs in the cache.
func (c *SyncCache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Len()
}

// Set the given key value pair.
// This operation updates the recent usage of the item.
func (c *SyncCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Set(key, value)
}

// Get an item from the cache.
// This operation updates recent usage of the item.
func (c *SyncCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache.Get(key)
}

// Peek gets an item from the cache without updating the recent usage.
func (c *SyncCache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Peek(key)
}

// Delete an item from the cache.
func (c *SyncCache[K, V]) Delete(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache.Delete(key)
}

// Flush deletes all items from the cache.
func (c *SyncCache[K, V]) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Flush()
}
