package main

import (
	"encoding/json"
	"log"
	"os"
)

var cfg = &Configuration{}

type Configuration struct {
	SrcMac string
	SrcIp  string
}

func LoadConfig(configFile string) {

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

}
