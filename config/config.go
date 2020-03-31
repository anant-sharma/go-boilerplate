package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Jwt Definition
type Jwt struct {
	Secret    string
	Algorithm string
	ExpiresIn time.Duration
	Issuer    string
}

// Configuration Definition
type Configuration struct {
	PORT               int
	DBConnectionString string
	Jwt                Jwt
}

/*
GetConfig - Return config
*/
func GetConfig() Configuration {
	loadConfig()

	appPort, _ := strconv.Atoi(os.Getenv("PORT"))
	jwtExpiresIn, _ := strconv.ParseInt(os.Getenv("jwt.expiresIn"), 10, 0)

	return Configuration{
		PORT:               appPort,
		DBConnectionString: os.Getenv("DBConnectionString"),
		Jwt: Jwt{
			Secret:    os.Getenv("jwt.secret"),
			Algorithm: os.Getenv("jwt.secret"),
			ExpiresIn: time.Duration(jwtExpiresIn),
			Issuer:    os.Getenv("jwt.issuer"),
		},
	}
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		os.Exit(500)
	}
}
