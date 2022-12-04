package behavioralPatterns

import "fmt"

// Lru is the least recently used
type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
