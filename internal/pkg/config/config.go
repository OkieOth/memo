package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

/*
The `Config` structure holds all configuration parameter
that are availe for the memo app
*/
type Config struct {
	// Directory to store the single files with the memos
	TargetDir string

	// Default target to use if no other target is given
	DefaultTarget string
}

func Get() Config {
	// look for $HOME/.memo/config.json
	conf, err := getFromFile("$HOME/.memo/config.json")
	if err != nil {
		conf = getDefaultConfig()
	}
	return conf
}

func getFromFile(filepath string) (Config, error) {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return Config{}, errors.New("config file doesn't exist")
	} else {
		if info.IsDir() {
			return Config{}, errors.New("given config file is a directory")
		}
	}
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()
	var config Config
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	return config, err
}

func getDefaultConfig() Config {
	conf := Config{}
	conf.TargetDir = "$HOME/.memo/targets"
	conf.DefaultTarget = "default"
	return conf
}
