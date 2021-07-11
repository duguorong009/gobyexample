package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "value 1"
	}()

	select {
	case msg := <-chan1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1s")
	}

	chan2 := make(chan int32, 1)
	go func() {
		time.Sleep(3 * time.Second)
		chan2 <- 444
	}()

	select {
	case number := <-chan2:
		fmt.Printf("Number is %d\n", number)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout 2")
	}

}
