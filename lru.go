package lru

import (
	"github.com/dboslee/lru/internal"
)

var _ internal.LRU[int, int] = &Cache[int, int]{}

// Cache is a lru cache. It automatically removes elements as new elements are
// added if the capacity is reached. Items are removes based on how recently
// they were used where the oldest items are removed first.
type Cache[K comparable, V any] struct {
	ll      *internal.List[entry[K, V]]
	items   map[K]*internal.Element[entry[K, V]]
	options *options
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// New initializes a new lru cache with the given capacity.
func New[K comparable, V any](cacheOptions ...CacheOption) *Cache[K, V] {
	c := &Cache[K, V]{
		ll:      internal.NewList[entry[K, V]](),
		items:   make(map[K]*internal.Element[entry[K, V]]),
		options: defaultOptions(),
	}

	for _, option := range cacheOptions {
		option.apply(c.options)
	}

	return c
}

// Len is the number of key value pairs in the cache.
func (c *Cache[K, V]) Len() int {
	return c.ll.Len()
}

// Set the given key value pair.
// This operation updates the recent usage of the item.
func (c *Cache[K, V]) Set(key K, value V) {
	if element, ok := c.items[key]; ok {
		element.Value.value = value
		c.ll.MoveToFront(element)
		return
	}

	entry := entry[K, V]{
		key:   key,
		value: value,
	}

	e := c.ll.PushFront(entry)
	if c.ll.Len() > c.options.capacity {
		c.deleteElement(c.ll.Back())
	}
	c.items[key] = e
}

// Get an item from the cache.
// This operation updates recent usage of the item.
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	e, ok := c.items[key]
	if !ok {
		return
	}

	c.ll.MoveToFront(e)
	return e.Value.value, true
}

// Peek gets an item from the cache without updating the recent usage.
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	e, ok := c.items[key]
	if !ok {
		return
	}

	return e.Value.value, true
}

// Delete an item from the cache.
func (c *Cache[K, V]) Delete(key K) bool {
	e, ok := c.items[key]
	if !ok {
		return false
	}

	c.deleteElement(e)

	return true
}

func (c *Cache[K, V]) deleteElement(e *internal.Element[entry[K, V]]) {
	delete(c.items, e.Value.key)
	c.ll.Remove(e)
}

// Flush deletes all items from the cache.
func (c *Cache[K, V]) Flush() {
	c.ll.Init()
	c.items = make(map[K]*internal.Element[entry[K, V]])
}
