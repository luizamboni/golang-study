package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type Cache struct {
	redis  *redis.Client
	inited bool
	prefix string
}

var instance Cache

func readConfig(relativePath string, filename string) viper.Viper {

	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName(filename)
	v.AddConfigPath(relativePath)
	v.AddConfigPath(".")
	err := v.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return *v
}

func GetInstance() (*Cache, error) {

	if instance.inited == false {
		config := readConfig("./config/", "redis")

		url := config.GetString("redis.url")
		password := config.GetString("redis.password")
		db := config.GetInt("redis.db")
		prefix := config.GetString("redis.prefix")

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

func (c Cache) Get(key string) string {

	str, _ := c.redis.Get(c.genKey(key)).Result()
	return str
}

func (c Cache) Set(key string, value string, seconds int) error {
	return c.redis.Set(c.genKey(key), value, time.Duration(seconds)*time.Second).Err()
}
