// Code from https://go101.org/article/concurrent-common-mistakes.html

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	// ruleid: waitgroup-wait-inside-loop
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)
			atomic.AddInt32(&x, 1)
		}()
		wg.Wait()
	}

	fmt.Println("Wait ...")
	fmt.Println(atomic.LoadInt32(&x))
}