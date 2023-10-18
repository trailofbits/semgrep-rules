package main

import (
	"fmt"
	"sync"
)

func main() {
	// In the body of the `for` add a call to one of the functions below to see
	// if a `fatal error: concurrent map writes` occurs
	// for i := 0; i < 1000; i++ {
	// 	fmt.Println("appendToMap1:")
	// 	res1 := appendToMap1()
	// 	printMap(res1)
	// }

	fmt.Println("appendToMap1:")
	res1 := appendToMap1()
	printMap(res1)

	fmt.Println("appendToMap2:")
	myMap := make(map[int]string)
	myMap[-1] = "number-negative-one"
	res2 := appendToMap2(myMap)
	printMap(res2)

	fmt.Println("appendToMap_FP1:")
	res3 := appendToMap_FP1()
	printMap(res3)

	fmt.Println("appendToMap_FP2:")
	res4 := appendToMap_FP2()
	printMap(res4)

	fmt.Println("appendToMap_FP3:")
	res5 := appendToMap_FP3()
	printMap(res5)

	fmt.Println("appendToMap_FP4:")
	res6 := appendToMap_FP4()
	printMap(res6)

	fmt.Println("appendToMap_FP5:")
	res7 := appendToMap_FP5()
	printMap(res7)
}

func appendToMap1() map[int]string {
	var wg sync.WaitGroup
	r := make(map[int]string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			// ruleid: racy-write-to-map
			r[iCpy] = fmt.Sprintf("number-%d", iCpy)
		}(i)
	}

	wg.Wait()

	return r
}

func appendToMap2(r map[int]string) map[int]string {
	var wg sync.WaitGroup
	r = make(map[int]string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			// ruleid: racy-write-to-map
			r[iCpy] = fmt.Sprintf("number-%d", iCpy)
		}(i)
	}

	wg.Wait()

	return r
}

func appendToMap_FP1() map[int]string {
	// FP: The map write is done inside a lock
	var wg sync.WaitGroup
	r := make(map[int]string)
	var rMut sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			rMut.Lock()
			// ok: racy-write-to-map
			r[iCpy] = fmt.Sprintf("number-%d", iCpy)
			rMut.Unlock()
		}(i)
	}

	wg.Wait()

	return r
}

func appendToMap_FP2() map[int]string {
	// FP: The map write is done in the body of the function and not in the goroutine
	var wg sync.WaitGroup
	r := make(map[int]string)

	for i := 0; i < 10; i++ {
		// ok: racy-write-to-map
		r[i] = fmt.Sprintf("number-%d", i)
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

func appendToMap_FP3() map[int]string {
	// FP: The map write inside the `for` loop of the main function and not in the goroutine
	var wg sync.WaitGroup
	r := make(map[int]string)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// ok: racy-write-to-map
			r[j] = fmt.Sprintf("number-%d", j)
		}
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	return r
}

func appendToMap_FP4() map[int]string {
	// FP: The map write happens on a map created in the goroutine
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			r := make(map[int]string)
			// ok: racy-write-to-map
			r[iCpy] = fmt.Sprintf("number-%d", iCpy)
		}(i)
	}

	wg.Wait()

	r := make(map[int]string)
	return r
}

func appendToMap_FP5() map[int]string {
	// FP: The map write is done inside a defer lock
	var wg sync.WaitGroup
	r := make(map[int]string)
	var rMut sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			rMut.Lock()
			fmt.Println("test")
			defer rMut.Unlock()
			// ok: racy-write-to-map
			r[iCpy] = fmt.Sprintf("number-%d", iCpy)
		}(i)
	}

	wg.Wait()

	return r
}

func printMap(results map[int]string) {
	for k, v := range results {
		fmt.Printf("key: %d, val: %s\n", k, v)
	}
}
