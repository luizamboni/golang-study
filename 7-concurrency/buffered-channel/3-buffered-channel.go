package main

import "fmt"

func main() {
	buffered := make(chan string, 2)

	buffered <- "text 1"
	buffered <- "text 2"

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

}
