package cyclemap

import (
	"cmp"
	"sync"
)

type CycleMap[K cmp.Ordered, V any] struct {
	keys   []K
	mapVal map[K]V
	index  int
	isSafe bool
	mu     *sync.Mutex
	size   int
}

func (c *CycleMap[K, V]) Get(k K) (V, bool) {
	if c.isSafe {
		c.mu.Lock()
		defer c.mu.Unlock()
	}
	v, ok := c.mapVal[k]
	return v, ok
}

func (c *CycleMap[K, V]) Set(k K, v V) {
	if c.isSafe {
		c.mu.Lock()
		defer c.mu.Unlock()
	}
	if _, ok := c.mapVal[k]; ok {
		c.mapVal[k] = v
		return
	}
	if len(c.keys) == c.size {
		c.index++
		if c.index >= c.size {
			c.index = 0
		}
		oldK := c.keys[c.index]
		delete(c.mapVal, oldK)
		c.keys[c.index] = k
		c.mapVal[k] = v
		return
	}
	c.index++
	c.keys = append(c.keys, k)
	c.mapVal[k] = v
}

func New[K cmp.Ordered, V any](size int, isSafe bool) *CycleMap[K, V] {

	return &CycleMap[K, V]{
		keys:   make([]K, 0, size),
		mapVal: make(map[K]V, size),
		index:  -1,
		isSafe: isSafe,
		mu:     &sync.Mutex{},
		size:   size,
	}
}
