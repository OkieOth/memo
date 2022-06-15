package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extractEnvVars(s string) []string {
	pattern := `\${?[[:alpha:]]+[[:alnum:]]*}?`
	re := regexp.MustCompile(pattern)
	result := re.FindAll([]byte(s), -1)
	if result == nil {
		return make([]string, 0)
	} else {
		tmp := make([]string, len(result))
		for i, b := range result {
			tmp[i] = string(b[:])
		}
		return tmp
	}
}

/*
Extracts all env vars from a string and replaces
them by the variable values from the system

example input:
IchBin${USER}AndMyHomeIs:${HOME}
$HOME/$UID/$GID
*/
func ReplaceEnvVars(s string) (string, error) {
	strRet := s
	pattern := `\${?([[:alpha:]]+[[:alnum:]]*)}?`
	re := regexp.MustCompile(pattern)
	varStrSlice := extractEnvVars(s)
	for _, varStr := range varStrSlice {
		matches := re.FindStringSubmatch(varStr)
		varName := matches[1]
		envVar := os.Getenv(varName)
		if envVar == "" {
			return s, errors.New(fmt.Sprintf("No value for %s found", varName))
		}
		strRet = strings.ReplaceAll(strRet, varStr, envVar)
	}
	return strRet, nil
}
