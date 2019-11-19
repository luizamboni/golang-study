package main

import (
	"fmt"
	"time"

	s "example.com/m/service"
	"github.com/go-redis/redis"
)

type Storage interface {
	Get(string) (string, error)
	Set(string, string, int) error
}

type Cache struct {
	redis  *redis.Client
	inited bool
	prefix string
}

var instance Cache

func GetInstance() (*Cache, error) {

	if instance.inited == false {
		url := "localhost:6379"
		password := ""
		db := 0
		prefix := "ooppa"

		redisclient := redis.NewClient(&redis.Options{
			Addr:     url,
			Password: password,
			DB:       db,
		})

		if err := redisclient.Ping().Err(); err != nil {
			return nil, err
		}

		instance = Cache{
			redis:  redisclient,
			inited: true,
			prefix: prefix,
		}
	}

	return &instance, nil
}
func (c Cache) genKey(key string) string {
	return c.prefix + key
}

func (c Cache) Get(key string) (string, error) {
	return c.redis.Get(c.genKey(key)).Result()
}

func (c Cache) Set(key string, value string, seconds int) error {
	return c.redis.Set(c.genKey(key), value, time.Duration(seconds)*time.Second).Err()
}

/* Service Example */
type ServiceX struct {
	c Storage
}

func (s ServiceX) getOfferByID(id string) (string, error) {
	fmt.Println(s, s.c)
	return s.c.Get("offer-" + id)
}

func main() {
	c, error := GetInstance()
	c, error = GetInstance()
	c, error = GetInstance()

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

	service := ServiceX{c}

	fmt.Println(service.c)

	c.Set("offer-abc", "123", 60)

	offer, error := service.getOfferByID("abc")

	if error != nil {
		fmt.Println("error", error)
	}

	fmt.Println("offer:", offer)

	fmt.Println(new(s.Service))
}
