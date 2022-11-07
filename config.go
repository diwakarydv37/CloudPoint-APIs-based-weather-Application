package config

import (
	"log"

	"github.com/diwakarydv37"
)

// Config of environment
type Config struct {
	Port string `env:"PORT" envDefault:"3000"`
}

// Get returns the environment configs
func Get() (cfg Config) {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return
}