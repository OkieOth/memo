package config

import (
	"os"
	"os/user"
	"testing"
)

func TestGetDefaultConfig(t *testing.T) {
	config := GetDefaultConfig()
	expected := "$HOME/.memo/targets"
	if config.TargetDir != expected {
		t.Errorf("Wrong default targetDir: %s, expected: %s", config.TargetDir, expected)
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
	config, err := GetFromFile("../../../configs/default_config.json")
	if err != nil {
		t.Errorf("received error while loading config: %v", err)
		return
	}
	if config.TargetDir != "~/.memo/targets" {
		t.Errorf("got wrong target dir in config: %s", config.TargetDir)
	}
}
