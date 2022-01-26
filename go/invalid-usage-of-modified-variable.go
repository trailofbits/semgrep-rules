package main

import (
	"fmt"
)

type Engineer struct {
	FName	string
	LName	string
	Age		int
	Address *Address
}

type Address struct {
	City	string
	State	string
}


func main() {
	engineers := getEngineers()
	// ruleid: invalid-usage-of-modified-variable
	eng1, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng1.FName)
	}

	// ruleid: invalid-usage-of-modified-variable
	if eng2, err := getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng2.FName)
	}

	eng3 := Engineer{"Lee", "Renaldo", 50, nil}
	// ruleid: invalid-usage-of-modified-variable
	eng3.Address, err = getEngineerAddressByIndex(engineers, 1)
	if err != nil {
		fmt.Printf("Unable to obtain address %#v!\n", eng3.Address)
	}

	fmt.Printf("Engineer 1: %s", fmt.Sprintf("%s %s", eng1.FName, eng1.LName))
}


func getEngineerAtIndex(slice []Engineer, idx int) (*Engineer, error) {
	if idx >= len(slice){
		return nil, fmt.Errorf("invalid index")
	}
	return &slice[idx], nil
}

func getEngineerAddressByIndex(slice []Engineer, idx int) (*Address, error) {
	if idx >= len(slice){
		return nil, fmt.Errorf("invalid index")
	}
	return slice[idx].Address, nil
}

func getEngineers() []Engineer {
	return []Engineer{
		{
			"James",
			"Osterberg",
			24,
			&Address {
				"Detroit",
				"Michigan",
			},
		},
		{
			"Althea",
			"Gordon",
			43,
			&Address {
				"New York",
				"New York",
			},
		},
		{
			"José-Manuel",
			"Arthur Chao",
			55,
			&Address {
				"Paris",
				"Île-de-France",
			},
		},
	}
}
