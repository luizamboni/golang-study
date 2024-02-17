package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	count := 0

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- 2
	}()

	for count < 2 {
		fmt.Println("waiting channels")
		// select is usad to wait multiple channels
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			count++

		case msg2 := <-c2:
			fmt.Println("received", msg2)
			count++
		}
	}
}
