package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var Server Configuration

func Load() {
	file, err := os.Open("./server.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	config, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(config, &Server)
	if err != nil {
		panic(err)
	}
}

type Configuration struct {
	Internal internal `json:"internal"`
	Database database `json:"database"`
}

type internal struct {
	Port                string        `json:"port"`
	APIKey              []string      `json:"apikey"`
	TokenExpirationTime time.Duration `json:"tokenExpirationTime"`
}

type database struct {
	SecondsToTimeOut int    `json:"timeOut"`
	InLocalDSN       string `json:"inLocalDsn"`
	InContainerDSN   string `json:"inContainerDsn"`
}
