package jdconf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Name    string `json:"name"`
	Threads int    `json:"threads"`
	Version string `json:"version"`
}

func readConfig(cfg *Config) error {

	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	configFileName = "/home/chan/Desktop/workspace/two/config.json"

	configFile, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		return err
	}
	return nil
}
