package main

import "sync"

type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
