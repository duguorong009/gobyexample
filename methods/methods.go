package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// when passing struct value
	r := rect{10, 20}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	// when passing struct pointer
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
