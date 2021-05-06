package main

import "sync"

type Counter struct {
	count int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
