package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProgs := os.Args
	argsWithoutProgs := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProgs)
	fmt.Println(argsWithoutProgs)
	fmt.Println(arg)
}
