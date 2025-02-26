package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Server struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Service struct {
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
}

type Config struct {
	Server   `yaml:"server"`
	Handlers map[string]Service `yaml:"handlers"`
}

func Get() *Config {
	return &config
}

var config Config

func init() {
	file, err := ioutil.ReadFile("./internal/config/config.yaml")
	if err != nil {
		log.Fatal("Error reading config file ", err)
		return
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error marshaling config ", err)
		return
	}

	log.Println("Config is loaded")
}
