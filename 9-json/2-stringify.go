package main

import (
	"encoding/json"
	"fmt"
)

type Complex struct {
	Name        string
	ExampleItem Example1
}
type Example1 struct {
	Id           int
	Key          string
	privateField string
}

func main() {
	bytes, err := json.Marshal(new(Complex))
	if err != nil {
		fmt.Println(err)
	}

	// private fields are out JSON
	fmt.Println("json:", string(bytes))
}
