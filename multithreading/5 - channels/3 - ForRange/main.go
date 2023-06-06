package main

import "fmt"

func main() {
	out := make(chan int)
	go reader(out)
	publish(out)
}

func publish(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func reader(out chan int) {
	for v := range out {
		fmt.Println(v)
	}
}
