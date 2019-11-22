package main

import "fmt"

type Storable interface {
	Set(string, string) bool
}

// StorageA
type StorageA struct{}

func (s StorageA) Set(key string, value string) bool {
	return true
}

func handleInfacePointer(st1 Storable, st2 Storable) {
	fmt.Println(st1.Set("a", "valor"))
	fmt.Println(st2.Set("a", "valor"))

}

func main() {

	st1 := StorageA{}
	handleInfacePointer(st1, &st1)
}
