package main

import (
	"fmt"
	"log"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	nextExample()
}

func nextExample() {
	log.Print("nextExample entry...")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan <- string, msg string) {
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}
