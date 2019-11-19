package main

import "fmt"

type String string

func (e String) Add(s String) {
	fmt.Println(e + " " + s)
}

func main() {
	s := String("A")
	s.Add(String("Adicionado"))
}
