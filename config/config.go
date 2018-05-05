package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

/*
Configuration Definition
*/
type Configuration struct {
	PORT               int
	DBConnectionString string
}

var configuration = Configuration{}

/*
GetConfig - Return config
*/
func GetConfig() Configuration {
	loadConfig()
	return configuration
}

func loadConfig() {

	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		log.Fatal("Err in config file", err)
		os.Exit(500)
	}

}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
