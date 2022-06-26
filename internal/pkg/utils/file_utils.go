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

func DoesDirExist(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false, nil
	} else {
		if !info.IsDir() {
			return false, errors.New("given directory is a file")
		}
	}
	return true, nil
}

/*
Creates a not existing directory.
The function returns a true if the directory was created or false if it alrady
existed. In case of errors the error is returned
*/
func CreateDirIfNotExist(dirPath string) (string, error) {
	path, err := ReplaceEnvVars(dirPath)
	if err != nil {
		return path, err
	}
	exists, err := DoesDirExist(path)
	if err != nil {
		return path, err
	}
	if exists {
		return path, nil
	}
	err = os.MkdirAll(path, os.ModePerm)
	if err == nil {
		return path, nil
	} else {
		return path, err
	}
}

func DeleteFileIfExist(filePath string) error {
	b, err := DoesFileExist(filePath)
	if err != nil {
		return err
	}
	if b {
		return os.Remove(filePath)
	}
	return nil
}
