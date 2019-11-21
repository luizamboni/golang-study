package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig(relativePath string, filename string) viper.Viper {

	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName(filename)
	v.AddConfigPath(relativePath)
	v.AddConfigPath(".")
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return *v
}
