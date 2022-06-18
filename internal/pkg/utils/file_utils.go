package utils

import (
	"errors"
	"os"
)

/**
Tests if a file exists the parameter can also contain env variables
that will be automatically replaced
*/
func DoesFileExist(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	} else {
		if info.IsDir() {
			return false, errors.New("given config file is a directory")
		}
	}
	return true, nil
}
