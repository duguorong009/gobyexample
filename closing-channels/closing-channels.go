package main

import "fmt"

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- "okok"
		fmt.Printf("sent %d job \n", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
