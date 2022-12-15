// Modified from https://gobyexample.com/mutexes

package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func main() {
	c := Container{
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

func (c *Container) inc(name string) error {
	c.mu.Lock()
	c.counters[name]++
	if name == "c" {
		// ruleid: missing-unlock-before-return
		return fmt.Errorf("letter not allowed")
	}
	c.mu.Unlock()
	return nil
}

func (c *Container) inc_FP(name string) error {
	c.mu.Lock()
	c.counters[name]++
	if name == "c" {
		// ok: missing-unlock-before-return
		c.mu.Unlock()
		return fmt.Errorf("letter not allowed")
	}
	c.mu.Unlock()
	return nil
}

// This could still cause a deadlock in the case that the caller function
// implements a recover() to avoid the application from crashing
func (c *Container) inc2(name string) error {
	c.mu.Lock()
	c.counters[name]++
	if name == "c" {
		// ruleid: missing-unlock-before-return
		panic("letter not allowed")
	}
	c.mu.Unlock()
	return nil
}

func (c *Container) inc2_FP(name string) error {
	c.mu.Lock()
	c.counters[name]++
	if name == "c" {
		c.mu.Unlock()
		// ok: missing-unlock-before-return
		panic("letter not allowed")
	}
	c.mu.Unlock()
	return nil
}

func (c *Container) inc3(name string) error {
	c.mu.Lock()
	c.counters[name]++
	if name == "c" {
		c.mu.Unlock()
		// ok: missing-unlock-before-return
		return fmt.Errorf("letter not allowed")
	}
	// ruleid: missing-unlock-before-return
	return nil
}
