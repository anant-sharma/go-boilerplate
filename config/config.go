package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Grpc struct {
		Host string
		Port int
	}
	Http struct {
		Host string
		Port int
	}
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

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Unable to load config to struct, %v", err)
	}
}
