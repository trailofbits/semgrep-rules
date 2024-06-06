
package main

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
)

func main() {
	iter1()
	iter1_FP_1()
	iter1_FP_2()
	iter1_FP_3()
	iter1_FP_4()
	iter1_FP_5()
}

func iter1(){
	// ruleid: iterate-over-empty-map
	m := make(map[string]int)

	fmt.Println("iterating")
	for v := range m {
		fmt.Println("map item: ", v)
	}
}

func iter1_FP_1(){
	// ok: iterate-over-empty-map
	m := make(map[string]int)

	m["v1"] = 8
	m["v2"] = 19

	for v := range m {
		fmt.Println("map item: ", v)
	}
}


func iter1_FP_2(){
	// ok: iterate-over-empty-map
	m := make(map[string]int)

	m["v1"]++

	for v := range m {
		fmt.Println("map item: ", v)
	}
}


func iter1_FP_3(){
	// ok: iterate-over-empty-map
	m := make(map[string]int)

	m["v1"]--

	for v := range m {
		fmt.Println("map item: ", v)
	}
}

func iter1_FP_4(){
	jsonValues := []byte(`{"one": 1, "two": 2, "three": 3}`)

	// ok: iterate-over-empty-map
	m := make(map[string]int)

	json.Unmarshal(jsonValues, &m)

	for v := range m {
		fmt.Println("map item: ", v)
	}
}

func iter1_FP_5() {
	// https://github.com/semgrep/semgrep/issues/9558
	// todook: iterate-over-empty-map
	testMap := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	keys := make([]string, 0, len(testMap))

	for k := range testMap {
		keys = append(keys, k)
	}
}
