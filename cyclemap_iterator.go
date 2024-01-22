package cyclemap

import "cmp"

type cyclemapIterator[K cmp.Ordered, V any] struct {
	cmap  *CycleMap[K, V]
	index int
}

func (c *cyclemapIterator[K, V]) First() (V, bool) {
	for c.index = 0; c.index < len(c.cmap.keys); c.index++ {
		key := c.cmap.keys[c.index]
		if v, ok := c.cmap.mapVal[key]; ok {
			return v, true
		}
	}
	var empty V
	return empty, false
}
func (c *cyclemapIterator[K, V]) Next() (V, bool) {
	for c.index++; c.index < len(c.cmap.keys); c.index++ {
		key := c.cmap.keys[c.index]
		if v, ok := c.cmap.mapVal[key]; ok {
			return v, true
		}
	}
	var empty V
	return empty, false
}
func (c *cyclemapIterator[K, V]) Index() int {
	return c.index
}
