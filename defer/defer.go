package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(fileName string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(fileName)
	if err != nil {
		panic("Error on creating file")
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("Writing...")
	fmt.Fprintf(f, "content is ...")
}
