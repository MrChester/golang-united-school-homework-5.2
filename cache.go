package cache

import (
	"time"
)

type Cache struct {
	kvPairs map[string]string
	keys    []string
}

func NewCache() Cache {
	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	_, ok := c.kvPairs[key]
	if !ok {
		return "", false
	}
	return c.kvPairs[key], true
}

func (c *Cache) Put(key, value string) {
	if c.kvPairs == nil {
		c.kvPairs = make(map[string]string)
		c.kvPairs[key] = value
	} else {
		c.kvPairs[key] = value
	}
}

func (c Cache) Keys() []string {
	for k := range c.kvPairs {
		c.keys = append(c.keys, k)
	}

	return c.keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.Put(key, value)

	if time.Now().Minute() >= deadline.Minute() {
		delete(c.kvPairs, key)
	}
}
