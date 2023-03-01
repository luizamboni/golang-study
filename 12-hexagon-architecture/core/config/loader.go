package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	payload map[string]interface{}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfig(path string) *Config {
	var payload map[string]interface{}
	data, err := os.ReadFile(path)
	check(err)
	err = json.Unmarshal(data, &payload)
	check(err)

	return &Config{
		payload: payload,
	}
}

func (s Config) GetAsString(keyName string) string {
	value := s.payload[keyName].(string)

	if strings.HasPrefix(value, "ENV:") {
		tokens := strings.SplitAfter(value, "ENV:")
		envValue := tokens[1]
		return os.Getenv(envValue)
	}
	return s.payload[keyName].(string)
}

func (s Config) GetAsInt(keyName string) int64 {
	value := s.payload[keyName].(string)

	if strings.HasPrefix(value, "ENV:") {
		tokens := strings.SplitAfter(value, "ENV:")
		envValue := tokens[1]
		intValue, err := strconv.ParseInt(envValue, 10, 64)
		check(err)
		return intValue
	}
	return s.payload[keyName].(int64)

}
