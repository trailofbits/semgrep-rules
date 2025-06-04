
package main

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
)

func main() {
	emptyMapIter()
	notEmptyMapIter()
	iter1()
	iter2()
	iter3()
	iter1_FP_1()
	iter1_FP_2()
	iter1_FP_3()
	iter1_FP_4()
	iter1_FP_5()
	iter1_FP_6()
	iter1_FP_7()
}

func emptyMapIter() {
  // ruleid: iterate-over-empty-map
  c := make(map[string]string)
  for k := range c {
    // do stuff
  }
}

func notEmptyMapIter() {
   // ok: iterate-over-empty-map
  c := map[string]string {
    "foo": "bar"
  }
  for k := range c {
    // do stuff
  }
}

func iter1(){
	// ruleid: iterate-over-empty-map
	m := make(map[string]int)

	fmt.Println("iterating")
	for v := range m {
		fmt.Println("map item: ", v)
	}
}

func iter2(){
	// https://github.com/semgrep/semgrep/issues/9558
	// todoruleid: iterate-over-empty-map
	m := make(map[string]int, 10)

	fmt.Println("iterating")
	for v := range m {
		fmt.Println("map item: ", v)
	}
}

func iter3(){
       // ruleid: iterate-over-empty-map
       m := map[string]int{}
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
	// ok: iterate-over-empty-map
	testMap := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	keys := make([]string, 0, len(testMap))

	for k := range testMap {
		keys = append(keys, k)
	}
}

func iter1_FP_6(){
       // ok: iterate-over-empty-map
       m := make(map[string]int, 1)

       m["v1"]--

       for v := range m {
               fmt.Println("map item: ", v)
       }
}

func iter1_FP_7(){
       // ok: iterate-over-empty-map
       m := map[string]int{"test": 1}
       for v := range m {
               fmt.Println("map item: ", v)
       }
}