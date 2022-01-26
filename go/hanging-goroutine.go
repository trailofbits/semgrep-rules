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
	req1(1)
	time.Sleep(time.Second * 5)
	fmt.Println(result)
	fmt.Println(runtime.NumGoroutine())
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

func test() string {
	time.Sleep(time.Second * 2)
	return "very important data"
}
