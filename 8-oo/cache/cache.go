package cache

import (
	"time"

	"github.com/go-redis/redis"
)

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
