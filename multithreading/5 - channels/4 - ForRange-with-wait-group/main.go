package main

import (
	"fmt"
	"sync"
)

func main() {
	out := make(chan int)
	in := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)
	go publish(out)
	go reader(out, &wg, in)
	fmt.Println(<-in)
	wg.Wait()
	fmt.Println("AQ")
}

func publish(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
}

func reader(out chan int, wg *sync.WaitGroup, in chan int) {
	in <- 12
	for v := range out {
		fmt.Println(v)
		wg.Done()
	}

}
