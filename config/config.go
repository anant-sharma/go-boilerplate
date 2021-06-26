package config

import (
	configUtils "github.com/anant-sharma/go-boilerplate-utils/config"
	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
)

type config struct {
	Grpc struct {
		Host string
		Port int
	}
	HTTP struct {
		Host string
		Port int
	}
	NewRelic newrelictracing.Config
}

// C - Config Instance
var C config

// Load - Loads config
func Load() {
	configUtils.Load(&C)
}
