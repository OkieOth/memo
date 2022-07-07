package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"okieoth/memo/internal/pkg/utils"
	"os"
	f "path/filepath"
)

const DEFAULT_CONFIG_PATH = "$HOME/.memo/config.json"

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
	conf, err := getFromFile(DEFAULT_CONFIG_PATH)
	if err != nil {
		conf = getDefaultConfig()
	}
	return conf
}

func (c Config) Write() error {
	return writeToFile("$HOME/.memo/config.json", c)
}

func writeToFile(filepath string, config Config) error {
	dir, _ := f.Split(filepath)
	_, err := utils.CreateDirIfNotExist(dir)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// TODO
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
