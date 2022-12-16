// Code from https://go101.org/article/concurrent-common-mistakes.html

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	// ruleid: waitgroup-wait-inside-loop
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&x, 1)
		}()
		wg.Wait()
	}

	fmt.Println("Wait ...")
	fmt.Println(atomic.LoadInt32(&x))
}

func test2() {
	// ruleid: waitgroup-wait-inside-loop
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("Wait ...")
	fmt.Println(atomic.LoadInt32(&x))
}


func test3() {
	// ok: waitgroup-wait-inside-loop
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&x, 1)
		}()
	}
	
	fmt.Println("Wait ...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}