package behavioralPatterns

// Strategy interface
type EvictionAlgo interface {
	evict(c *Cache)
}
