package main

import "fmt"

func varfunc(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total = total + num
	}
	fmt.Println(total)
}

func main() {
	varfunc(1, 2)

	varfunc(1, 2, 3)

	slice := []int{1, 2, 3}
	varfunc(slice...)
}
