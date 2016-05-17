package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type DBServerConfig struct {

	ServerLocation string `json:"serverLocation"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`

}

// Gets the server configuration from the configuration file
func GetDBServerConfig() DBServerConfig {

	configFile, err := os.Open("config/db_config.json")
	if err != nil {
        fmt.Println("opening config file: ", err.Error())
    }

    jsonParser := json.NewDecoder(configFile)

    config := DBServerConfig{}
    if err := jsonParser.Decode(&config); err != nil {
    	fmt.Println("decoding config file: ", err.Error())
    }

    return config

}