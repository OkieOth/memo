package utils

import (
	"regexp"
)

/*
Extracts all env vars from a string and replaces
them by the variable values from the system

example input:
IchBin${USER}AndMyHomeIs:${HOME}
$HOME/$UID/$GID
*/
func ReplaceEnvVars(s string) []string {
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
