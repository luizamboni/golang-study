package main

import (
	"fmt"
	"strconv"
)

func printSomething(i int, channel chan string) {
	fmt.Println("start of printSomething", i)

	channel <- "okokoko " + strconv.FormatInt(int64(i), 10)
	fmt.Println("end of printSomething", i)
}

func main() {
	n := make([]int, 10)
	channel := make(chan string)

	// the program exit before all gorroutines executes
	for i := range n {
		go printSomething(i, channel)
	}

	// each <- pops one value
	for _ = range n {
		fmt.Println(<-channel)
	}
}
