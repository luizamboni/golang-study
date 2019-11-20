package cache

import (
	"fmt"
	"io/ioutil"
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

func readFile(relativePath string) (string, error) {
	b, err := ioutil.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}
	return string(b), nil
}

func GetInstance() (*Cache, error) {

	if instance.inited == false {
		// content, _ := readFile("./config/redis.toml")
		viper.SetConfigType("toml")
		viper.SetConfigName("redis")
		viper.AddConfigPath("./config/")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		url := viper.GetString("redis.url")
		password := viper.GetString("redis.password")
		db := viper.GetInt("redis.db")
		prefix := viper.GetString("redis.prefix")

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
