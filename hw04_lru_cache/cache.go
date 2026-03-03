package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	v, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(v)
		c.items[key].Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		last := c.queue.Back()

		for key, v := range c.items {
			if v == last {
				delete(c.items, key)
				break
			}
		}

		c.queue.Remove(last)
	}

	c.queue.PushFront(value)
	c.items[key] = c.queue.Front()

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	v, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(v)
	return v.Value, ok
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
