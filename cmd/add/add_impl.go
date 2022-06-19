package add

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"okieoth/memo/internal/pkg/config"
	"regexp"
	"strings"
)

type Memo struct {
	Text    string
	Targets []string
}

func IsInteractiveMode(args *[]string) bool {
	if len(*args) == 0 {
		return true
	}
	return false
}

func GetMemoFromStdin(stdin io.Reader) (Memo, error) {
	var memo Memo
	scanner := bufio.NewScanner(stdin)
	fmt.Print("Enter the memo: ")
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		return memo, errors.New("no input, so further processing is skipped ... bye")
	}
	splittedText := strings.Fields(text)
	return ParseInput(&splittedText)
}

func ParseInput(inputStrings *[]string) (Memo, error) {
	var memo Memo
	r, _ := regexp.Compile("#[a-zA-Z0-9.-_]*")
	for _, s := range *inputStrings {
		if r.MatchString(s) {
			memo.Targets = append(memo.Targets, s[1:])
		} else {
			if len(memo.Text) == 0 {
				memo.Text = s
			} else {
				memo.Text = memo.Text + " " + s
			}
		}
	}
	if memo.Text == "" {
		return memo, errors.New("no text given, nothing to proccess ... bye")
	} else {
		return memo, nil
	}
}

func storeNow(target string, text string, config config.Config) {
	//targetDir := config.TargetDir

}

func StoreMemo(memo Memo, config config.Config) error {
	if len(memo.Targets) == 0 {
		// store the memo in the default target
		storeNow(config.DefaultTarget, memo.Text, config)
	} else {
		for _, t := range memo.Targets {
			storeNow(t, memo.Text, config)
		}
	}
	return errors.New("TODO")
}
