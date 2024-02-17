package main

import "fmt"

func main() {
	messages := make(chan string)
	// if try in same goroutive, wilt generate a error
	// uncomment to see
	// messages <- "test"
	go func() {
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}
