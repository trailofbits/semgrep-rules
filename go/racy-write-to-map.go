package main

import (
	"fmt"
	"sync"
)

func main() {
	res1 := appendToMap1()
	printMap(res1)

	myMap :=  make(map[int]string)
	myMap[-1] = "number-negative-one"
	res2 := appendToMap2(myMap)
	printMap(res2)
}

func appendToMap1() map[int]string {
	var wg sync.WaitGroup
	r := make(map[int]string)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(iCpy int) {
			defer wg.Done()
			// ruleid: racy-write-to-map
			r[i] = fmt.Sprintf("number-%d", i)
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
			r[i] = fmt.Sprintf("number-%d", i)
		}(i)
	}

	wg.Wait()

	return r
}

func printMap(results map[int]string) {
	for k, v := range results {
		fmt.Printf("key: %d, val: %s", k, v)
	}
}