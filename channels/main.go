package main

import (
	"fmt"
	"log"
)

func main() {
	//1
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	//2
	nextExample()

	//3
	c := wibblePostfixer("test", "test2", "test3")
	out := blahBlahPrefixer(c)

	for o := range out {
		fmt.Println(o)
	}

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
	msg := <-pings
	pongs <- msg
}

func wibblePostfixer(strings ...string) <- chan string {
	out := make(chan string)
	go func() {
		for _, n := range strings {
			out <- n + " wibble"
		}
		close(out)
	}()
	return out
}

func blahBlahPrefixer(in <- chan string) <- chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- "blah blah " + n
		}
		close(out)
	}()
	return out
}
