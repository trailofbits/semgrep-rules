package main

type TestStruct1 struct {
	// ok: unmarshal-tag-is-dash
	A string `json:"id"`
}

type TestStruct2 struct {
	// ruleid: unmarshal-tag-is-dash
	B string `json:"-,omitempty"`
}

type TestStruct3 struct {
	// ruleid: unmarshal-tag-is-dash
	C string `json:"-,123"`
}

type TestStruct4 struct {
	// ruleid: unmarshal-tag-is-dash
	D string `json:"-,"`
}
