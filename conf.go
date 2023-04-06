package main

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
)

//go:embed conf/config.yaml
var configData embed.FS

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func main() {
	configBytes, err := configData.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Database configuration: %+v\n", config.Database.Port)
}
