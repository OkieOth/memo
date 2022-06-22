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
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	file2 := fmt.Sprintf("%s/../../../main.go", path)
	doesExist, err = DoesFileExist(file2) // call the function with a real file
	if !doesExist {
		t.Errorf("file was not found: %s", file2)
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDoesDirExist(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error("Error while query current working directory")
		return
	}
	exists, e := DoesDirExist(path)
	if (!exists) || (e != nil) {
		t.Errorf("DoesDirExist couldn't find working dir: exists=%v, err=%v, path=%s", exists, e, path)
	}
	file1 := fmt.Sprintf("%s/env_var.go", path)
	exists, e = DoesDirExist(file1) // call the function with a real file
	if (exists) || (e == nil) {
		t.Errorf("DoesDirExist couldn't find file as dir: exists=%v, err=%v, path=%s", exists, e, file1)
	}
	exists, e = DoesDirExist("/this/is/a/fake/dir")
	if (exists) || (e != nil) {
		t.Errorf("DoesDirExist couldn't detect a not existing dir: exists=%v, err=%v, path=%s", exists, e, path)
	}
}

func TestCreateDirIfNotExist(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error("Error while query current working directory")
		return
	}
	pathStr, e := CreateDirIfNotExist(path)
	if (pathStr != path) || (e != nil) {
		t.Errorf("CreateDirIfNotExist couldn't find working dir: created=%v, err=%v, path=%s", pathStr, e, path)
	}

	dir1 := fmt.Sprintf("%s/tmp/TestCreateDirIfNotExist", path)
	pathStr, e = CreateDirIfNotExist(dir1)
	if (pathStr != dir1) || (e != nil) {
		t.Errorf("CreateDirIfNotExist didn't work: created=%v, err=%v, path=%s", pathStr, e, dir1)
	}
	created, e2 := DoesDirExist(dir1)
	if (!created) || (e2 != nil) {
		t.Errorf("DoesDirExist didn't find fresh created dir: created=%v, err=%v, path=%s", created, e2, dir1)
	}
	e2 = os.RemoveAll(dir1)
	if e2 != nil {
		t.Errorf("Couldn't remove fresh created dir: err=%v, path=%s", e2, dir1)
	}
	created, e2 = DoesDirExist(dir1)
	if (created) || (e2 != nil) {
		t.Errorf("DoesDirExist found fresh created dir: created=%v, err=%v, path=%s", created, e2, dir1)
	}
}
