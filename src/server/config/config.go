package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func Load() *Configuration {
	file, err := os.Open("./config-server.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Configuration

	err = json.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	return &config
}

type Configuration struct {
	Internal internal `json:"internal"`
	Database database `json:"database"`
	Cache    cache    `json:"cache"`
}

type internal struct {
	Port                string        `json:"port"`
	APIKey              []string      `json:"apiKey"`
	TokenExpirationTime time.Duration `json:"tokenExpirationTime"`
	EmailRegex          string        `json:"emailRegex"`
}

type database struct {
	SecondsToTimeOut time.Duration `json:"timeOut"`
	InLocalDSN       string        `json:"inLocalDsn"`
	InContainerDSN   string        `json:"inContainerDsn"`
}

type cache struct {
	LocalAddr     string `json:"localAddr"`
	ContainerAddr string `json:"containerAddr"`
}
