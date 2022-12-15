// Modified from https://gobyexample.com/mutexes

package main

import (
	"fmt"
	"sync"
)

type RWContainer struct {
	mu       sync.RWMutex
	counters map[string]int
}

func main() {
	c := RWContainer{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			err := c.inc(name)
			if err != nil {
				continue
			}
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("c", 10000)
	go doIncrement("a", 10000)
	go doIncrement("c", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}

func (c *RWContainer) inc(name string) error {
	c.mu.RLock()
	c.counters[name]++
	if name == "c" {
		// ruleid: missing-runlock-on-rwmutex
		return fmt.Errorf("letter not allowed")
	}
	c.mu.RUnlock()
	return nil
}

func (c *RWContainer) inc2(name string) error {
	c.mu.RLock()
	c.counters[name]++
	if name == "c" {
		// ruleid: missing-runlock-on-rwmutex
		return fmt.Errorf("letter not allowed")
	}
	// ruleid: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc_FP(name string) error {
	c.mu.RLock()
	c.counters[name]++
	if name == "c" {
		// ok: missing-runlock-on-rwmutex
		c.mu.RUnlock()
		return fmt.Errorf("letter not allowed")
	}
	c.mu.RUnlock()
	return nil
}
