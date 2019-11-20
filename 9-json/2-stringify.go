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
	Id  int
	Key string
}

func main() {
	bytes, err := json.Marshal(new(Complex))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json:", string(bytes))
}
