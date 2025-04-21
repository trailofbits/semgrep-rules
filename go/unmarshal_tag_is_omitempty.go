package main

import (
	"encoding/json"
	"fmt"
)

type SomeStruct struct {
	// ruleid: unmarshal-tag-is-omitempty
	SomeField string `json:"omitempty"`
	// ok: unmarshal-tag-is-omitempty
	SomeField_2 string `json:",omitempty"`
}

func test_omitempty() {
	u := SomeStruct{}
	_ = json.Unmarshal([]byte(`{"omitempty": "123", "SomeField_2": "456"}`), &u)
	fmt.Printf("Result: %#v\n", u)
	// Result: main.SomeStruct{SomeField:"123"}
}
