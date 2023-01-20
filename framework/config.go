package main

import (
	"fmt"
	"os"
)

var (
	systemConfigUrl = "../config/system.json"
	configValues    = make(map[string]string)
)

func loadConfig() {
	getLogger().Info("Loading configuration file at ", systemConfigUrl)
	jsonFile, err := os.Open(systemConfigUrl)
	if err != nil {
		fmt.Println(err)
		jsonFile.Close()
	}
}

func getValue(key string) string {
	return configValues[key]
}

func setValue() {

}
