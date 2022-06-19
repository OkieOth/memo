package config

import (
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
