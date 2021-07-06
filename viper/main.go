package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	cfg := viper.New()
	cfg.SetConfigFile("cfg.yml")
	cfg.ReadInConfig()
	fmt.Println(cfg.GetString("name"))
}
