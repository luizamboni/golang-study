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

// is like Example, but Id field is mapped to "identification" field
type Example2 struct {
	Id  int `json:"identification"`
	Key string
}

func main() {

	// simple Example1
	// singular object
	bytes := []byte(`{"id": 1,"key": "ab"}`)
	singular := new(Example1)
	json.Unmarshal(bytes, singular)
	fmt.Println("simple example:", *singular)

	// in Example2
	bytes2 := []byte(`{"identification": 1,"key": "ab"}`)
	singular2 := new(Example2)
	json.Unmarshal(bytes2, singular2)
	fmt.Println("example custom mapped:", *singular2)

	// plural
	bytes3 := []byte(`[{"id": 1,"key": "ab"},{"id": 2,"key": "cd"},{"id": 3,"key": "ef"}]`)
	collection := make([]Example1, 0)
	json.Unmarshal(bytes3, &collection)
	fmt.Println("collection:", collection)
}
