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

	res3 := appendToSlice_FP1()
	printResults(res3)

	res4 := appendToSlice_FP2()
	printResults(res4)

	res5 := appendToSlice_FP3()
	printResults(res5)

	res6 := appendToSlice_FP4()
	printResults(res6)

	res7 := appendToSlice_FP5()
	printResults(res7)
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

func appendToSlice_FP1() []int {
	// FP: The `append` is done inside a lock
	var wg sync.WaitGroup
	var rMut sync.Mutex
	var r []int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			m := iCpy * 2
			rMut.Lock()
			// ok: racy-append-to-slice
			r = append(r, m)
			rMut.Unlock()
		}(i)
	}

	wg.Wait()

	return r
}

func appendToSlice_FP2() []int {
	// FP: The `append` is done in the body of the function and not in the goroutine
	var wg sync.WaitGroup
	var r []int

	for i := 0; i < 10; i++ {
		// ok: racy-append-to-slice
		r = append(r, i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return r
}

func appendToSlice_FP3() []int {
	// FP: The `append` inside the `for` loop of the main function and not in the goroutine
	var wg sync.WaitGroup
	var r []int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; i++ {
			// ok: racy-append-to-slice
			r = append(r, j)
		}
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return r
}

func appendToSlice_FP4() []int {
	// FP: The `append` happens on a slice created in the goroutine
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			var r []int
			m := iCpy * 2
			// ok: racy-append-to-slice
			r = append(r, m)
		}(i)
	}

	wg.Wait()

	var r []int
	return r
}

func appendToSlice_FP5() []int {
	// FP: The `append` is done inside a defer lock
	var wg sync.WaitGroup
	var rMut sync.Mutex
	var r []int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			m := iCpy * 2
			rMut.Lock()
			fmt.Println("test")
			defer rMut.Unlock()
			// ok: racy-append-to-slice
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
