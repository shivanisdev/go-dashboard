package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Config  of the Server
type Config struct {
	DNS    string `json:"dns"`
	DbName string `json:"dbname"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

//LoadConfiguration  from the local file
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
