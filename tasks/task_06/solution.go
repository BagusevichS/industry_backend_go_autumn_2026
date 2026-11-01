package main

import "container/list"

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	order    list.List
	index    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		index:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	el, ok := c.index[key]
	if !ok {
		var zero V
		return zero, false
	}
	c.order.MoveToFront(el)
	return el.Value.(entry[K, V]).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	if el, ok := c.index[key]; ok {
		el.Value = entry[K, V]{key, value}
		return
	}
	if len(c.index) == c.capacity {
		oldest := c.order.Back()
		delete(c.index, oldest.Value.(entry[K, V]).key)
		c.order.Remove(oldest)
	}
	c.index[key] = c.order.PushFront(entry[K, V]{key, value})
}
