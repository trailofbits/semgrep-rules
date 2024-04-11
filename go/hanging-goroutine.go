package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	result = ""
)

func main() {
	req6(1, false)
	fmt.Println("Result: ", result)
	fmt.Println("Goroutines (must be 1 for FPs):", runtime.NumGoroutine())
}

func req1(timeout time.Duration) string {
	ch := make(chan string)
	// ruleid: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	select {
	case result = <- ch:
		fmt.Println("case result")
		return result
	case <- time.After(timeout):
		fmt.Println("case time.Afer")
		return ""
	}
}

func req1_FP(timeout time.Duration) string {
	ch := make(chan string, 1)
	// ok: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	select {
	case result = <- ch:
		fmt.Println("case result")
		return result
	case <- time.After(timeout):
		fmt.Println("case time.Afer")
		return ""
	}
}

func req2(timeout time.Duration) string {
	ch := make(chan string)
	// ruleid: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	select {
	case <- ch:
		fmt.Println("case result")
		return result
	case <- time.After(timeout):
		fmt.Println("case time.Afer")
		return ""
	}
}

func req2_FP(timeout time.Duration) string {
	ch := make(chan string, 1)
	// ok: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	select {
	case <- ch:
		fmt.Println("case result")
		return result
	case <- time.After(timeout):
		fmt.Println("case time.Afer")
		return ""
	}
}

func req3(timeout time.Duration) {
	ch := make(chan string)
	// ruleid: hanging-goroutine
	for i := 0; i < 3; i++ {
		go func() {
			newData := test()
			ch <- newData // block
		}()
	}
	result = <- ch
	fmt.Println("finished req3")
}

func req3_FP(timeout time.Duration) {
	ch := make(chan string, 3)
	// ok: hanging-goroutine
	for i := 0; i < 3; i++ {
		go func() {
			newData := test()
			ch <- newData // block
		}()
	}
	result = <- ch
	fmt.Println("finished req3")
}

func req4_FP(timeout time.Duration) string {
	ch := make(chan string)
	// ok: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	select {
	case result = <- ch:
		fmt.Println("case result")
		return result
	case <- time.After(timeout):
		result = <- ch
		fmt.Println("case time.Afer")
		return ""
	}
}

func req5_FP(timeout time.Duration) {
	ch := make(chan string)
	tick := time.Tick(100 * time.Millisecond)
	quit := time.After(2 * time.Second)
	// ok: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	for {
		select {
			case <-tick:
				fmt.Print("|")
			case <-quit:
				result = <- ch
				fmt.Println("\nquit")
				return
			default:
				fmt.Print(".")
				time.Sleep(50 * time.Millisecond)
		}
	}
}

func req6(timeout time.Duration, doclose bool) {
	ch := make(chan string)
	tick := time.Tick(100 * time.Millisecond)
	quit := time.After(2 * time.Second)
	// todoruleid: hanging-goroutine
	go func() {
		newData := test()
		ch <- newData // block
	}()
	for {
		select {
			case <-tick:
				fmt.Print("|")
			case <-quit:
				if doclose {
					result = <- ch
				}
				fmt.Println("\nquit")
				return
			default:
				fmt.Print(".")
				time.Sleep(50 * time.Millisecond)
		}
	}
}


func test() string {
	time.Sleep(time.Second * 2)
	return "very important data"
}
