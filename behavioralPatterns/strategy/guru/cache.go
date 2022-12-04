package behavioralPatterns

import "sync"

type Cache struct {
	rw           sync.Mutex
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	delete(c.storage, key)
}

func (c *Cache) evict() {
	if c.rw.TryLock() {
		defer c.rw.Unlock()
	}
	c.evictionAlgo.evict(c)
	c.capacity--
}
