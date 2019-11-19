package main

import (
	"fmt"

	"example.com/m/cache"

	"example.com/m/service"
)

func main() {
	c, error := cache.GetInstance()

	if error != nil {
		panic(error)
	}

	v, error := c.Get("abc")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("value:", v)

	c.Set("abc", "123", 6)

	v, error = c.Get("abc")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("value:", v)

	service := service.New(c)

	c.Set("offer-abc", "123", 60)

	offer, error := service.GetOfferByID("abc")

	if error != nil {
		fmt.Println("error", error)
	}

	fmt.Println("offer:", offer)

}
