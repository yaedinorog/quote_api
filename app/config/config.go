package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

var Cfg *Config

func LoadConfig(path_to_config string) {
	file, err := os.Open(path_to_config)

	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	Cfg = &Config{}
	if err := decoder.Decode(Cfg); err != nil {
		log.Fatalf("Error while reading from config file: %v", err)
	}
}
