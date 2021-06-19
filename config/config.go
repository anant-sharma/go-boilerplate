package config

import (
	"strings"

	newrelictracing "github.com/anant-sharma/go-utils/new-relic/tracing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	// Set the path to look for the configurations file
	viper.AddConfigPath("config")

	// Set the file name of the configurations file
	viper.SetConfigName("")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Convert _ underscore in env to . dot notation in viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Unable to load config to struct, %v", err)
	}
}
