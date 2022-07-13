package https

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	RedirectHost string   `json:"redirectHost"`
	HostsPolicy  []string `bson:"hostsPolicy"`
	AcmeEmail    string   `bson:"acmeEmail"`
}

func LoadConfiguration() (*Configuration, error) {
	var config Configuration
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			println(err.Error())
			return
		}
	}(configFile)
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
