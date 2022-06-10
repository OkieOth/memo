package add

import (
	"errors"
	"regexp"
)

type Memo struct {
	text    string
	targets []string
}

func IsInteractiveMode(args *[]string) bool {
	if len(*args) == 0 {
		return true
	}
	return false
}

func ParseInput(inputStrings *[]string) (Memo, error) {
	var memo Memo
	r, _ := regexp.Compile("#[a-zA-Z0-9.-_]*")
	for _, s := range *inputStrings {
		if r.MatchString(s) {
			memo.targets = append(memo.targets, s[1:])
		} else {
			if len(memo.text) == 0 {
				memo.text = s
			} else {
				memo.text = memo.text + " " + s
			}
		}
	}
	if memo.text == "" {
		return memo, errors.New("no text given")
	} else {
		return memo, nil
	}
}
