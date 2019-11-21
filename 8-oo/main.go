package main

import (
	"fmt"

	"example.com/m/service"

	"example.com/m/cache"
)

func main() {
	c, error := cache.GetInstance()

	if error != nil {
		panic(error)
	}

	v := c.Get("abc")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("value when not exists:", v)

	c.Set("abc", "{ \"Name\":\"notebook\", \"Count\": 123, \"Average\": 12, \"Price\": 100.0 }", 6)

	v = c.Get("abc")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("value:", v)

	service := service.New(c)

	c.Set("offer-1", "{ \"Name\": \"telefone\", \"Count\": 123, \"Average\": 12, \"Price\": 100.0 }", 60)
	c.Set("offer-2", "{ \"name\": \"telefone\", \"Count\": 123, \"Average\": 12, \"Price\": 100.0 }", 60)
	c.Set("offer-2", "{ \"NAME\": \"telefone\", \"Count\": 123, \"Average\": 12, \"Price\": 100.0 }", 60)

	offer1 := service.GetOfferByID("1")
	offer2 := service.GetOfferByID("1")
	offer3 := service.GetOfferByID("1")

	// if error != nil {
	// 	fmt.Println("error", error)
	// }

	fmt.Println("offer:", offer1)
	fmt.Println("offer:", offer2)
	fmt.Println("offer:", offer3)

}
