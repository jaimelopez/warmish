package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Sitemaps    []string
	Concurrency int
}

func (config *Config) Parse(filename string) error {
	data, error := ioutil.ReadFile(filename)

	if error != nil {
		return error
	}

	return yaml.Unmarshal(data, config)
}

func New(filename string) *Config {
	config := new(Config)
	error := config.Parse(filename)

	if error != nil {
		panic(error)
	}

	return config
}
