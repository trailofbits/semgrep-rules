package main

import (
	"fmt"
	"sync"
)

func main() {
	res1 := appendToSlice1()
	printResults(res1)

	res2 := appendToSlice2()
	printResults(res2)
}

func appendToSlice1() []int {
	var wg sync.WaitGroup
	var r []int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			m := iCpy * 2
			// ruleid: racy-append-to-slice
			r = append(r, m)
		}(i)
	}

	wg.Wait()

	return r
}

func appendToSlice2() []int {
	var wg sync.WaitGroup
	r := make([]int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			m := iCpy * 2
			// ruleid: racy-append-to-slice
			r = append(r, m)
		}(i)
	}

	wg.Wait()

	return r
}

func printResults(results []int) {
	for v, _ := range results {
		fmt.Println(v)
	}
}