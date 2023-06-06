package main

import "time"

type Message struct {
	id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	//RabbitMQ
	go func() {
		for {
			time.Sleep(time.Second * 2)
			msg := Message{1, "Hello from RabbitMQ"}
			c1 <- msg
		}

	}()

	//Kafka
	go func() {
		for {
			time.Sleep(time.Second * 1)
			msg := Message{1, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			println("received", msg.Msg)
		case msg := <-c2:
			println("received", msg.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
