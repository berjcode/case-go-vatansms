package config

import (
    "encoding/json"
"os"
)

type Configuration struct {
    Database DatabaseConfig `json:"Database"`
}

type DatabaseConfig struct {
    Username string `json:"Username"`
    Password string `json:"Password"`
    Host     string `json:"Host"`
    Port     string `json:"Port"`
    Name     string `json:"Name"`
}


func LoadConfiguration(filename string) (Configuration, error) {
    var config Configuration
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        return config, err
    }
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    return config, err
}