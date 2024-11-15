// Modified from https://gobyexample.com/mutexes

package main

import (
	"fmt"
	"sync"
)

type Unlocker func()

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

type Fridge struct {
	food int
}

func (f *Fridge) Lock() {
	fmt.Println(f.food)
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
	// ok: missing-unlock-before-return
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
	// ok: missing-unlock-before-return
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
	// ok: missing-unlock-before-return
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
	// ok: missing-unlock-before-return
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

func (c *Container) inc4FP(name string) Unlocker {
	c.mu.Lock()
	c.counters[name]++
	return func() {
		// ok: missing-unlock-before-return
		c.mu.Unlock()
	}
}

func (c *Container) inc5(name string) error {
	f := Fridge{food: 11}
	f.Lock()
	// ok: missing-unlock-before-return
	return nil
}

func (c *Container) inc6(name string) error {
	c.mu.Lock()
	c.counters[name]++
	defer func() {
		fmt.Println("before unlock")
		c.mu.Unlock()
		fmt.Println("after unlock")
	}()
	// ok: missing-unlock-before-return
	return nil
}

func (c *Container) inc6b(name string) error {
	c.mu.Lock()
	unlocker := c.mu.Unlock
	c.counters[name]++
	defer func() {
		fmt.Println("before unlock")
		unlocker()
		// todook: missing-unlock-before-return
		fmt.Println("after unlock")
	}()
	// todook: missing-unlock-before-return
	return nil
}

func (c *Container) inc7(name string) error {
	c.mu.Lock()
	c.counters[name]++
	defer func(earlyExit bool) {
		fmt.Println("before unlock")
		if (earlyExit) {
			// todoruleid: missing-unlock-before-return
			return
		}
		c.mu.Unlock()
		fmt.Println("after unlock")
	}(false)
	// ok: missing-unlock-before-return
	return nil
}

func (c *Container) inc8(name string) error {
	c.mu.Lock()
	c.counters[name]++
	_, err := fmt.Println("test if")
	if err != nil {
	    c.mu.Unlock()
	    // ok: missing-unlock-before-return
	    return nil, err
	}
	// ruleid: missing-unlock-before-return
	return nil
}

func (c *Container) inc8b(name string) error {
	c.mu.Lock()
	c.counters[name]++
	unlocker := c.mu.Unlock
	_, err := fmt.Println("test if")
	if err != nil {
	    unlocker()
	    // todook: missing-unlock-before-return
	    return nil, err
	}
	// ruleid: missing-unlock-before-return
	return nil
}
