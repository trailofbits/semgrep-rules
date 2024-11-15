// Modified from https://gobyexample.com/mutexes

package main

import (
	"fmt"
	"sync"
)

type RUnlocker func()

type RWContainer struct {
	mu       sync.RWMutex
	counters map[string]int
}

type Fridge struct {
	food int
}

func (f *Fridge) RLock() {
	fmt.Println(f.food)
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
	// ok: missing-runlock-on-rwmutex
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
	// ok: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc4FP(name string) RUnlocker {
	c.mu.Lock()
	c.counters[name]++
	return func() {
		// ok: missing-runlock-on-rwmutex
		c.mu.Unlock()
	}
}

func (c *RWContainer) inc5(name string) error {
	f := Fridge{food: 11}
	f.RLock()
	// ok: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc6(name string) error {
	c.mu.RLock()
	c.counters[name]++
	defer func() {
		fmt.Println("before runlock")
		c.mu.RUnlock()
		fmt.Println("after runlock")
	}()
	// ok: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc6b(name string) error {
	c.mu.RLock()
	unlocker := c.mu.RUnlock
	c.counters[name]++
	defer func() {
		fmt.Println("before runlock")
		unlocker()
		// todook: missing-runlock-on-rwmutex
		fmt.Println("after runlock")
	}()
	// todook: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc7(name string) error {
	c.mu.RLock()
	c.counters[name]++
	defer func(earlyExit bool) {
		fmt.Println("before runlock")
		if (earlyExit) {
			// todoruleid: missing-runlock-on-rwmutex
			return
		}
		c.mu.RUnlock()
		fmt.Println("after runlock")
	}(false)
	// ok: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc8(name string) error {
	c.mu.RLock()
	c.counters[name]++
	_, err := fmt.Println("test if")
	if err != nil {
	    c.mu.RUnlock()
	    // ok: missing-runlock-on-rwmutex
	    return nil, err
	}
	// ruleid: missing-runlock-on-rwmutex
	return nil
}

func (c *RWContainer) inc8b(name string) error {
	c.mu.RLock()
	c.counters[name]++
	unlocker := c.mu.RUnlock
	_, err := fmt.Println("test if")
	if err != nil {
	    unlocker()
	    // todook: missing-runlock-on-rwmutex
	    return nil, err
	}
	// ruleid: missing-runlock-on-rwmutex
	return nil
}
