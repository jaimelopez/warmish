package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	DEFAULT_PURGE       = true
	DEFAULT_WARMUP      = false
	DEFAULT_CONCURRENCY = 10
	DEFAULT_BREAK       = "0ms"
)

type Config struct {
	Sitemaps    []string
	Purge       bool
	Warmup      bool
	Concurrency uint
	Break       string
	Schedule    string
}

func (config *Config) Parse(filename string) error {
	data, error := ioutil.ReadFile(filename)

	if error != nil {
		return error
	}

	return yaml.Unmarshal(data, config)
}

func New(filename string) *Config {
	config := &Config{
		Purge:       DEFAULT_PURGE,
		Warmup:      DEFAULT_WARMUP,
		Concurrency: DEFAULT_CONCURRENCY,
		Break:       DEFAULT_BREAK,
	}

	error := config.Parse(filename)

	if error != nil {
		panic(error)
	}

	return config
}
