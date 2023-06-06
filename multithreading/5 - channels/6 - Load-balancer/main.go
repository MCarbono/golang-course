package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data <-chan int) {
	for v := range data {
		fmt.Printf("Worker %d received %d\n", workerID, v)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	workers := 100000
	for i := 0; i < workers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
