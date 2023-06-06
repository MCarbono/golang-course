package main

import "fmt"

func main() {
	ch := make(chan string)
	go receive("Hello", ch)
	read(ch)
}

func receive(name string, ch chan<- string) {
	ch <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}
