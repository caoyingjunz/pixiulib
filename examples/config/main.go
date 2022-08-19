package main

import (
	"fmt"

	"github.com/caoyingjunz/pixiulib/config"
)

type Config struct {
	Default DefaultOption `yaml:"default"`
}

type DefaultOption struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	c := config.New()
	c.SetConfigFile("./config.yaml")
	c.SetConfigType("yaml")

	var cfg Config
	if err := c.Binding(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
