// Code from https://go101.org/article/concurrent-common-mistakes.html

package main

import "sync"

type Counter struct {
	sync.Mutex
	n int64
}


// ok: sync-mutex-value-copied
func (c *Counter) Increase(d int64) (r int64) {
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

// ruleid: sync-mutex-value-copied
func (c Counter) Value() (r int64) {
	c.Lock()
	r = c.n
	c.Unlock()
	return
}
