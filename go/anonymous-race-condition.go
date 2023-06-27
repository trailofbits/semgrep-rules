package main

import (
	"fmt"
	"sync"
)

type myval struct {
	s string
}
type myval2 struct {
	s myval
}

func (v *myval) Method(x string) {
	fmt.Println(v.s + "-" + x)
}

var (
	numbers = [6]string{"one", "two", "three", "four", "five", "six"}
)

var (
	values = [3]myval2{myval2{myval{"a"}}, myval2{myval{"b"}}, myval2{myval{"c"}}}
)

func main() {
	ConcurrentFunctions_TP(func1, func2)
	ConcurrentFunctions_FP(func1, func2)
	AnonRaceCond_1()
	AnonRaceCond_1_FP()
	AnonRaceCond_2()
	AnonRaceCond_2_FP()
	AnonRaceCond_3_FP()
	AnonRaceCond_4()
	AnonRaceCond_4_FP()
}

func func1() {
	fmt.Println("I am function func1")
}

func func2() {
	fmt.Println("I am function func2")
}


func ConcurrentFunctions_TP(fns ...func()) {
	var wg sync.WaitGroup
	// ruleid: anonymous-race-condition
	for _, fn := range fns {
		wg.Add(1)
		go func() {
			fn()
			wg.Done()
		}()
	}

	wg.Wait()
}

func ConcurrentFunctions_FP(fns ...func()) {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for _, fn := range fns {
		wg.Add(1)
		go func(fn2 func()) {
			fn2()
			wg.Done()
		}(fn)
	}

	wg.Wait()
}

func AnonRaceCond_1() {
	var wg sync.WaitGroup
	// ruleid: anonymous-race-condition
	for _, num := range numbers {
		wg.Add(1)
		go func() {
			fmt.Println(num)
			wg.Done()
		}()
	}

	wg.Wait()
}

func AnonRaceCond_1_FP() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for _, num := range numbers {
		wg.Add(1)
		go func(cpy string) {
			fmt.Println(cpy)
			wg.Done()
		}(num)
	}

	wg.Wait()
}

func AnonRaceCond_2() {
	var wg sync.WaitGroup
	// ruleid: anonymous-race-condition
	for i := 0; i< len(numbers); i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}


func AnonRaceCond_2_FP() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for i := 0; i< len(numbers); i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func AnonRaceCond_3_FP() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for i := 0; i< len(numbers); i++ {
		cpy := i
		wg.Add(1)
		go func() {
			fmt.Println(cpy)
			wg.Done()
		}()
	}

	wg.Wait()
}

func AnonRaceCond_4() {
	var wg sync.WaitGroup
	// ruleid: anonymous-race-condition
	for _, val := range values {
		wg.Add(1)
		go func() {
			val.s.Method("test")
			wg.Done()
		}()
	}

	wg.Wait()
}

func AnonRaceCond_4_FP() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for _, val := range values {
		wg.Add(1)
		val := val
		go func() {
			val.s.Method("test")
			wg.Done()
		}()
	}

	wg.Wait()
}


func AnonRaceCond_5() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for idx, val := range values {
		wg.Add(1)
		idx, val := idx, val
		go func() {
			val.s.Method("test")
			fmt.Println("Completed index", idx)
			wg.Done()
		}()
	}

	wg.Wait()
}

func AnonRaceCond_6() {
	var wg sync.WaitGroup
	// ok: anonymous-race-condition
	for idx, val := range values {
		wg.Add(1)
		idx, val := idx, val
		go func() {
			fmt.Println(val.s)
			fmt.Println("Completed index", idx)
			wg.Done()
		}()
	}

	wg.Wait()
}
