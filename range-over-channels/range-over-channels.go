package main

import "fmt"

func main() {
	queue := make(chan string, 4)

	queue <- "start"
	queue <- "realy?"

	close(queue)

	for elem := range queue {
		fmt.Printf("Element: %s \n", elem)
	}
}
