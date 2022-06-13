//nolint:typecheck
package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/luisnquin/go-redoc"
)

func Load() *Configuration {
	file, err := os.Open("./config-server.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var config Configuration

	if err = json.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}

	return &config
}

//nolint:typecheck
type Configuration struct {
	Internal internal    `json:"internal"`
	Database database    `json:"database"`
	Cache    cache       `json:"cache"`
	Docs     redoc.Redoc `json:"docs"`
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
	SchemasPath      string        `json:"schemasPath"`
}

type cache struct {
	LocalAddr     string `json:"localAddr"`
	ContainerAddr string `json:"containerAddr"`
}
