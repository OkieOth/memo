package add

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"okieoth/memo/internal/pkg/config"
	"okieoth/memo/internal/pkg/utils"
	"os"
	"regexp"
	"strings"
	"time"
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

func storeNow(target string, text string, config config.Config) error {
	targetDir, err := utils.CreateDirIfNotExist(config.TargetDir)
	if err != nil {
		return err
	}
	targetFileName := fmt.Sprintf("%s/%s.md", targetDir, target)
	f, err := os.OpenFile(targetFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

func StoreMemo(memo Memo, config config.Config) error {
	timestamp := time.Now()
	nowStr := timestamp.Format("20060102140405")
	outputTxt := fmt.Sprintf("* %s [%s]\n", memo.Text, nowStr)
	if len(memo.Targets) == 0 {
		// store the memo in the default target
		err := storeNow(config.DefaultTarget, outputTxt, config)
		if err == nil {
			fmt.Printf("  created memo: %s-%s\n", config.DefaultTarget, nowStr)
		}
		return err
	} else {
		for _, t := range memo.Targets {
			err := storeNow(t, outputTxt, config)
			if err == nil {
				return err
			} else {
				fmt.Printf("  created memo: %s-%s", t, nowStr)
			}
		}
	}
	return nil
}
