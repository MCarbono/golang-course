package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		c <- "Hello World!"
	}()

	msg := <-c
	fmt.Println(msg)
}
