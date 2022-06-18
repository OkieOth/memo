package utils

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestDoesFileExist(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error("Error while query current working directory")
		return
	}
	pattern := `memo/internal/pkg/utils$`
	re := regexp.MustCompile(pattern)
	if !re.Match([]byte(path)) {
		t.Error("Error while query current working directory")
		return
	}
	var doesExist bool
	doesExist, err = DoesFileExist(path) // call the function with the current working dir
	if doesExist {
		t.Errorf("directory detected as file: %s", path)
		return
	}
	if err == nil {
		t.Errorf("no error returned when tested workdir as file: %s", path)
		return
	}
	file1 := fmt.Sprintf("%s/env_var.go", path)
	doesExist, err = DoesFileExist(file1) // call the function with a real file
	if !doesExist {
		t.Errorf("file was not found: %s", file1)
		return
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	file2 := fmt.Sprintf("%s/../../../main.go", path)
	doesExist, err = DoesFileExist(file2) // call the function with a real file
	if !doesExist {
		t.Errorf("file was not found: %s", file2)
		return
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
}
