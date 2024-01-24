package main

import (
	"fmt"
)

type Engineer struct {
	Id      int
	FName   string
	LName   string
	Age     int
	Address *Address
}

func (e *Engineer) String() string {
	return e.FName + " " + e.LName
}

type Address struct {
	City  string
	State string
}

type Log struct {
	Id      int
	Message string
}

func main() {
	engineers := getEngineers()

	// BASIC TESTS
	// ruleid: invalid-usage-of-modified-variable
	eng1, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng1.FName)
	}

	// ruleid: invalid-usage-of-modified-variable
	eng11, err2 := getEngineerAtIndex(engineers, 5)
	if err2 != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng11.FName)
		fmt.Printf("Something")
	}

	var eng12 *Engineer
	// ruleid: invalid-usage-of-modified-variable
	eng12, err = getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng12.FName)
	}

	var eng13 *Engineer
	// ruleid: invalid-usage-of-modified-variable
	eng13, err = getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng13.FName)
		fmt.Printf("Something")
	}

	// ruleid: invalid-usage-of-modified-variable
	eng14, addr1, err := getEngineerAndAddressByIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng14.FName)
		fmt.Printf("Something")
	}

	// BASIC TESTS - inline if
	// ruleid: invalid-usage-of-modified-variable
	if eng2, err := getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng2.FName)
	}

	// ruleid: invalid-usage-of-modified-variable
	if eng21, err := getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng21.FName)
		fmt.Printf("Something")
	}

	var eng22 *Engineer
	// ruleid: invalid-usage-of-modified-variable
	if eng22, err = getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng22.FName)
	}

	var eng23 *Engineer
	// ruleid: invalid-usage-of-modified-variable
	if eng23, err = getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng23.FName)
		fmt.Printf("Something")
	}

	var eng24 *Engineer
	// ruleid: invalid-usage-of-modified-variable
	if eng24, err = getEngineerAtIndex(engineers, 6); err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer with FName %s\n", eng24.FName)
		fmt.Printf("Something")
		eng24 = &Engineer{0, "N/A", "N/A", 0, nil}
	}

	// WHOLE VAR AND METHODS TESTS
	// ruleid: invalid-usage-of-modified-variable
	eng5, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("%s\n", eng5.String())
		fmt.Printf("Something")
	}

	// FALSE POSITIVES
	eng3 := Engineer{4, "Lee", "Renaldo", 50, nil}
	// ok: invalid-usage-of-modified-variable
	eng3.Address, err = getEngineerAddressByIndex(engineers, 1)
	if err != nil {
		fmt.Printf("Unable to obtain address %#v!\n", eng3.Address)
	}

	// ok: invalid-usage-of-modified-variable
	eng3.Address, err = getEngineerAddressByIndex(engineers, 1)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain address %#v!\n", eng3.Address)
		fmt.Printf("Something")
	}

	// ok: invalid-usage-of-modified-variable
	eng4, err := getEngineerAtIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("%#v\n", eng4)
		fmt.Printf("Something")
	}

	// ok: invalid-usage-of-modified-variable
	eng6, err := getEngineerAtIndex(engineers, 7)
	if err != nil {
		eng6 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng6.Id)
	}

	// ok: invalid-usage-of-modified-variable
	eng61, err := getEngineerAtIndex(engineers, 7)
	if err != nil {
		fmt.Printf("Something")
		eng61 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng61.Id)
		fmt.Printf("Something")
	}

	// ok: invalid-usage-of-modified-variable
	eng7, err := getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng7 := &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng7.Id)
	}

	// ok: invalid-usage-of-modified-variable
	eng71, err := getEngineerAtIndex(engineers, 8)
	if err != nil {
		fmt.Printf("Something")
		eng71 := &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng71.Id)
		fmt.Printf("Something")
	}

	var eng8 *Engineer
	// ok: invalid-usage-of-modified-variable
	eng8, err = getEngineerAtIndex(engineers, 7)
	if err != nil {
		eng8 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng8.Id)
	}

	var eng81 *Engineer
	// ok: invalid-usage-of-modified-variable
	eng81, err = getEngineerAtIndex(engineers, 7)
	if err != nil {
		fmt.Printf("Something")
		eng81 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng81.Id)
		fmt.Printf("Something")
	}

	// ok: invalid-usage-of-modified-variable
	eng9, err := getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng9 = &Engineer{0, "N/A", "N/A", 0, nil}
	}

	var eng91 *Engineer
	// ok: invalid-usage-of-modified-variable
	eng91, err = getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng91 = &Engineer{0, "N/A", "N/A", 0, nil}
	}

	// ok: invalid-usage-of-modified-variable
	eng92, addr2, err := getEngineerAndAddressByIndex(engineers, 5)
	if err != nil {
		fmt.Printf("Something")
		eng92, err = getEngineerAtIndex(engineers, 6)
		fmt.Printf("Something")
	}

	// ok: invalid-usage-of-modified-variable
	eng93, err := getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng93 = &Engineer{0, "N/A", "N/A", 0, nil}
		eng93 = &Engineer{eng93.Id, "N/A", "N/A", 0, nil}
	}

	// ok: invalid-usage-of-modified-variable
	eng94, err := getEngineerAtOrDefault(engineers, 8)
	if err != nil {
		if eng94 != nil {
			fmt.Printf("Got engineer %s %s\n", eng94.FName, eng94.LName)
		}
	}

	// ok: invalid-usage-of-modified-variable
	eng95, err := getEngineerAtOrDefault(engineers, 8)
	if err != nil && eng95.FName == "Doe" {
		fmt.Println("Got a generic engineer")
	}

	// TRUE POSITIVES

	// ruleid: invalid-usage-of-modified-variable
	eng11, err = getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng11.Address = nil
	}

	// ruleid: invalid-usage-of-modified-variable
	eng12, err = getEngineerAtIndex(engineers, 8)
	if err != nil {
		eng12 = &Engineer{eng12.Id, "N/A", "N/A", 0, nil}
	}

	// The following test case should match, but I was unable to find a way to
	// match it without causing some of the false positives to match. This is
	// largely due to the requirement that "<... $X.$Y ...>" be the last line in
	// the if-block, i.e. the following causes an invalid pattern error:
	//  - pattern: |
	//      $X, err = ...
	//      if err != nil {
	//          ...
	//          <... $X.$Y ...>
	//          ...
	//          $X = ...
	//      }

	// todoruleid: invalid-usage-of-modified-variable
	eng10, err := getEngineerAtIndex(engineers, 10)
	if err != nil {
		fmt.Printf("Something")
		fmt.Printf("Unable to obtain engineer %d!\n", eng10.Id)
		eng10 = &Engineer{0, "N/A", "N/A", 0, nil}
		fmt.Printf("Unable to obtain engineer %d!\n", eng10.Id)
	}

	fmt.Printf("Engineer 1: %s", fmt.Sprintf("%s %s", eng1.FName, eng1.LName))
	fmt.Printf("Engineer 7: %s", fmt.Sprintf("%s %s", eng7.FName, eng7.LName))
	fmt.Printf("Engineer 71: %s", fmt.Sprintf("%s %s", eng7.FName, eng71.LName))
	fmt.Printf("Engineer 9: %s", fmt.Sprintf("%s %s", eng9.FName, eng9.LName))
	fmt.Printf("Engineer 91: %s", fmt.Sprintf("%s %s", eng91.FName, eng91.LName))
	fmt.Printf("Engineer 92: %s", fmt.Sprintf("%s %s", eng92.FName, eng92.LName))
	fmt.Printf("Address 1: %v", addr1)
	fmt.Printf("Address 2: %v", addr2)
}

func getEngineerAtIndex(slice []Engineer, idx int) (*Engineer, error) {
	if idx >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return &slice[idx], nil
}

func getEngineerAddressByIndex(slice []Engineer, idx int) (*Address, error) {
	if idx >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return slice[idx].Address, nil
}

func getEngineerAndAddressByIndex(slice []Engineer, idx int) (*Engineer, *Address, error) {
	if idx >= len(slice) {
		return nil, nil, fmt.Errorf("invalid index")
	}
	return &slice[idx], slice[idx].Address, nil
}

func getEngineerAtOrDefault(slice []Engineer, idx int) (*Engineer, error) {
	if idx >= len(slice) {
		return nil, fmt.Errorf("invalid index")
	}
	return &slice[idx], nil
}

func defaultEngineer() Engineer {
	return Engineer{
		1,
		"Jane",
		"Doe",
		24,
		&Address{
			"San Marcos",
			"California",
		},
	}
}

func getEngineers() []Engineer {
	return []Engineer{
		{
			1,
			"James",
			"Osterberg",
			24,
			&Address{
				"Detroit",
				"Michigan",
			},
		},
		{
			2,
			"Althea",
			"Gordon",
			43,
			&Address{
				"New York",
				"New York",
			},
		},
		{
			3,
			"José-Manuel",
			"Arthur Chao",
			55,
			&Address{
				"Paris",
				"Île-de-France",
			},
		},
	}
}
