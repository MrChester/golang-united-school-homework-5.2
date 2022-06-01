package cache

import (
	"fmt"
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
	fmt.Println(c)
	_, ok := c.kvPairs[key]
	if !ok {
		return "", false
	}
	return c.kvPairs[key], true
}

func (c *Cache) Put(key, value string) {
	c.keys = append(c.keys, key)
	if c.kvPairs == nil {
		c.kvPairs = make(map[string]string)
		c.kvPairs[key] = value
	} else {
		c.kvPairs[key] = value
	}

	fmt.Println("Here we putting the value")
}

func (c Cache) Keys() []string {
	return c.keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	fmt.Println(deadline)
}
