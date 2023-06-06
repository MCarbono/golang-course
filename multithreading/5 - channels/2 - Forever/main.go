package main

func main() {
	forever := make(chan bool)
	forever <- true
	<-forever
}
