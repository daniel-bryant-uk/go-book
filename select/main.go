package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
		wg.Done()
	}()

	wg.Wait()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
