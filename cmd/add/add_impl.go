package add

import (
	"bufio"
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
	Header  string
}

func getStringFromStdin(stdin io.Reader, inputMsg string, canBeEmpty bool) string {
	scanner := bufio.NewScanner(stdin)
	fmt.Print(inputMsg)
	scanner.Scan()
	text := scanner.Text()
	for (len(text) == 0) && (!canBeEmpty) {
		fmt.Println(inputMsg)
		scanner.Scan()
		text = scanner.Text()
	}
	return text
}

func getMemoTextFromStdin(stdin io.Reader) string {
	return getStringFromStdin(stdin, "Enter a text (cancel with CTRL-C): ", false)
}

func getMemoHeaderFromStdin(stdin io.Reader) string {
	return getStringFromStdin(stdin, "Enter a Header (cancel with CTRL-C): ", true)
}

func getTargetsFromStdin(stdin io.Reader) string {
	r, _ := regexp.Compile("[a-zA-Z0-9.-_ ]+")
	scanner := bufio.NewScanner(stdin)
	fmt.Print("Enter targets separated by space: ")
	scanner.Scan()
	text := scanner.Text()
	for len(text) == 0 {
		fmt.Println("Enter a targets or cancel with CTRL-C")
		scanner.Scan()
		text = scanner.Text()
		if len(text) > 0 {
			if !r.MatchString(text) {
				fmt.Println(" ...wrong input! Only [a-zA-Z0-9.-_] allowed.")
				text = ""
			}
		}
	}
	return text
}

type InitStdin struct {
	Stdin     io.Reader
	MockStdin [3]io.Reader
}

type StdinInputSource int64

const (
	Text StdinInputSource = iota
	Target
	Header
)

func (stdinInput InitStdin) get(inputSource StdinInputSource) io.Reader {
	if stdinInput.Stdin != nil {
		return stdinInput.Stdin
	}
	switch inputSource {
	case Text:
		return (stdinInput.MockStdin)[0]
	case Target:
		return (stdinInput.MockStdin)[1]
	case Header:
		return (stdinInput.MockStdin)[2]
	default:
		panic(fmt.Sprintf("Wrong input source type: %v", inputSource))
	}
}

/*
Init the memo parameters from stdin.
*/
func InitFromStdin(stdin InitStdin) (string, string, string) {
	text := getMemoTextFromStdin(stdin.get(Text))
	targets := getTargetsFromStdin(stdin.get(Target))
	header := getMemoHeaderFromStdin(stdin.get(Header))
	return text, targets, header
}

func InitMemoFromStdin(stdin InitStdin, memo *Memo) {
	if memo == nil {
		panic("Memo is nil! There is nothing to init from stdin")
	}

	text, targets, header := InitFromStdin(stdin)
	memo.Text = text
	memo.Targets = strings.Fields(targets)
	memo.Header = header
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
			fmt.Printf("  store memo: %s-%s\n", nowStr, config.DefaultTarget)
		}
		return err
	} else {
		for _, t := range memo.Targets {
			err := storeNow(t, outputTxt, config)
			if err != nil {
				return err
			} else {
				fmt.Printf("  store memo: %s-%s\n", nowStr, t)
			}
		}
	}
	return nil
}
