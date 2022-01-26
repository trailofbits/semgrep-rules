// code modified from https://gobyexample.com/methods

package main

import "fmt"

type rect struct {
	width, height int
}

func main() {
	rect1()
	rect2()
	rect3()
	rect4()
}

func NewRectangle(w, h int) (*rect, error) {
	if w == 0 || h == 0 {
		return nil, fmt.Errorf("invalid height or width ")
	}
	return &rect{width: w, height: h}, nil
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) print() {
	fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
}

func rect1(){
	r, _ := NewRectangle(0, 2)
	// ruleid: nil-check-after-call
	fmt.Println("area: ", r.area())
	if r != nil {
		fmt.Println("perim:", r.perim())
	}
}

func rect2(){
	r, _ := NewRectangle(0, 2)
	// ruleid: nil-check-after-call
	r.print()
	if r != nil {
		fmt.Println("perim:", r.perim())
	}
}

func rect3(){
	r, err := NewRectangle(0, 2)
	if err != nil {
		return
	}
	r.print()
	r, err = NewRectangle(2, 3)
	// ruleid: nil-check-after-call
	r.print()
	if r != nil {
		fmt.Println("perim:", r.perim())
	}
}

func rect4(){
	r, err := NewRectangle(0, 2)
	if err != nil {
		return
	}
	r.print()

	r, err = NewRectangle(2, 3)
	// ruleid: nil-check-after-call
	fmt.Println("perim:", r.perim())
	if r == nil {
		return
	}
}