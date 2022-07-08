package config

import (
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Printf("no config file found (%s), go with the default config ...\n\n", DEFAULT_CONFIG_PATH)
		conf = getDefaultConfig()
	}
	return conf
}

func (c Config) Write() error {
	return writeToFile("$HOME/.memo/config.json", &c)
}

func (c Config) AsJson() ([]byte, error) {
	jsonContent, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return make([]byte, 0), err
	}
	return jsonContent, nil
}

func writeToFile(filepath string, config *Config) error {
	dir, file := f.Split(filepath)
	absDir, err := utils.CreateDirIfNotExist(dir)
	if err != nil {
		return err
	}

	jsonContent, err := config.AsJson()
	if err != nil {
		return err
	}

	fp := fmt.Sprintf("%s%c%s", absDir, os.PathSeparator, file)
	return ioutil.WriteFile(fp, jsonContent, 0644)
}

func getFromFile(filepath string) (Config, error) {
	f, _ := utils.ReplaceEnvVars(filepath)
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return Config{}, errors.New("config file doesn't exist")
	} else {
		if info.IsDir() {
			return Config{}, errors.New("given config file is a directory")
		}
	}
	jsonFile, err := os.Open(f)
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
