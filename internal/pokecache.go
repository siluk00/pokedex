package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]CacheEntry
	mutex sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(dur time.Duration) *Cache {
	c := Cache{data: make(map[string]CacheEntry)}
	c.ReapLoop(dur)
	return &c
}

func (c *Cache) Add(key string, value []byte) {
	cacheEntry := CacheEntry{createdAt: time.Now(), val: value}
	c.mutex.Lock()
	c.data[key] = cacheEntry
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, ok := c.data[key]
	c.mutex.Unlock()

	if ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) ReapLoop(dur time.Duration) {
	tickerDuration, err := time.ParseDuration("5s")
	if err != nil {
		fmt.Printf("Error parsing duration: %s", err)
	}

	ticker := time.NewTicker(tickerDuration)
	go func() {
		for t := range ticker.C {
			c.mutex.Lock()
			for k, v := range c.data {
				if t.Compare(v.createdAt.Add(dur)) > 0 {
					delete(c.data, k)
				}
			}
			c.mutex.Unlock()
		}
	}()
}
