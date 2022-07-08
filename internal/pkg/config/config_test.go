package config

import (
	"okieoth/memo/internal/pkg/utils"
	"os"
	"os/user"
	"testing"
)

func TestGetDefaultConfig(t *testing.T) {
	config := getDefaultConfig()
	expected := "$HOME/.memo/targets"
	if config.TargetDir != expected {
		t.Errorf("Wrong default targetDir: %s, expected: %s", config.TargetDir, expected)
	}
	expected = "default"
	if config.DefaultTarget != expected {
		t.Errorf("Wrong default target: %s, expected: %s", config.DefaultTarget, expected)
	}
}

func TestOpenFile(t *testing.T) {
	user, _ := user.Current()
	filepath := user.HomeDir
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		t.Errorf("Can't open home dir: %v", err)
	}
	if !info.IsDir() {
		t.Error("Home dir is not directory?")
	}
}

func TestGetFromFile(t *testing.T) {
	config, err := getFromFile("../../../configs/default_config.json")
	if err != nil {
		t.Errorf("received error while loading config: %v", err)
		return
	}
	if config.TargetDir != "~/.memo/targets" {
		t.Errorf("got wrong target dir in config: %s", config.TargetDir)
	}
	if config.DefaultTarget != "other_stuff" {
		t.Errorf("got wrong default target in config: %s", config.DefaultTarget)
	}
	config, err = getFromFile("../../../configs/default_config_xxx.json")
	if err == nil {
		t.Error("didn't receive an error when reading config from not existing file")
		return
	}
	config, err = getFromFile("../../../configs")
	if err == nil {
		t.Error("didn't receive an error when reading directory")
		return
	}
	config, err = getFromFile("config_test.go")
	if err == nil {
		t.Error("didn't receive an error when reading non-json file")
		return
	}
}

func TestGet(t *testing.T) {
	conf := Get()
	if conf.DefaultTarget == "" {
		t.Error("got empty defaultTarget")
	}
	if conf.TargetDir == "" {
		t.Error("got empty targetDir")
	}
}

func TestWrite(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Error("Error while query current working directory")
		return
	}
	config, err := getFromFile("../../../configs/default_config.json")
	if err != nil {
		t.Errorf("received error while loading config: %v", err)
		return
	}
	if config.TargetDir != "~/.memo/targets" {
		t.Errorf("got wrong target dir in config: %s", config.TargetDir)
	}
	if config.DefaultTarget != "other_stuff" {
		t.Errorf("got wrong default target in config: %s", config.DefaultTarget)
	}
	destFile := workingDir + "/../../tmp/TestWrite.json"
	if writeToFile(destFile, &config) != nil {
		t.Errorf("error while store the config: %s", destFile)
	}
	config2, err2 := getFromFile(destFile)
	if err2 != nil {
		t.Errorf("received error while loading config (2): %v", err2)
		return
	}
	if config != config2 {
		t.Errorf("loaded config differs compared to the original: \noriginal: %v\nloaded: %v", config, config2)
	}
	config2.DefaultTarget = "default"
	config2.TargetDir = "/tmp"
	if config == config2 {
		t.Errorf("config wasn't changed: \noriginal: %v\nchanged: %v", config, config2)
	}
	if writeToFile(destFile, &config2) != nil {
		t.Errorf("error while store the config (2): %s", destFile)
	}
	config3, err3 := getFromFile(destFile)
	if err3 != nil {
		t.Errorf("received error while loading config (2): %v", err3)
		return
	}
	if config2 != config3 {
		t.Errorf("seems that changed file wasn't right written: \noriginal: %v\nloaded: %v", config2, config3)
	}
	_ = utils.DeleteFileIfExist(destFile)
}
