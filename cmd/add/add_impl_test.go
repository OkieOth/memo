package add

import (
	"bytes"
	"okieoth/memo/internal/pkg/config"
	"okieoth/memo/internal/pkg/utils"
	"os"
	"testing"
)

func TestInitMemoFromStdin_1(t *testing.T) {
	var stdinMock1 bytes.Buffer
	var stdinMock2 bytes.Buffer
	var stdinMock3 bytes.Buffer

	stdinMock1.Write([]byte("This is the memo text\n"))
	stdinMock2.Write([]byte("target1 target2 target3\n"))
	stdinMock3.Write([]byte("I am a header\n"))
	var memo Memo

	var stdin InitStdin
	stdin.MockStdin[0] = &stdinMock1
	stdin.MockStdin[1] = &stdinMock2
	stdin.MockStdin[2] = &stdinMock3

	InitMemoFromStdin(stdin, &memo)
	if memo.Text != "This is the memo text" {
		t.Errorf("Wrong Text: %s", memo.Text)
	}
	if (len(memo.Targets) != 3) || (memo.Targets[0] != "target1") ||
		(memo.Targets[1] != "target2") || (memo.Targets[2] != "target3") {
		t.Errorf("Wrong Targets: %v", memo.Targets)
	}
	if memo.Header != "I am a header" {
		t.Errorf("Wrong Header: %s", memo.Header)
	}
}

func TestInitMemoFromStdin_2(t *testing.T) {
	var stdinMock1 bytes.Buffer
	var stdinMock2 bytes.Buffer
	var stdinMock3 bytes.Buffer

	stdinMock1.Write([]byte("This is the memo text\n"))
	stdinMock2.Write([]byte("target1\n"))
	stdinMock3.Write([]byte("\n"))
	var memo Memo

	var stdin InitStdin
	stdin.MockStdin[0] = &stdinMock1
	stdin.MockStdin[1] = &stdinMock2
	stdin.MockStdin[2] = &stdinMock3

	InitMemoFromStdin(stdin, &memo)
	if memo.Text != "This is the memo text" {
		t.Errorf("Wrong Text: %s", memo.Text)
	}
	if (len(memo.Targets) != 1) || (memo.Targets[0] != "target1") {
		t.Errorf("Wrong Targets: %v", memo.Targets)
	}
	if memo.Header != "" {
		t.Errorf("Wrong Header: %s", memo.Header)
	}
}

func TestMemo(t *testing.T) {
	memo1 := Memo{}
	if memo1.Text != "" {
		t.Errorf(`memo1.text isn't "", instead %v`, memo1.Text)
	}
	if memo1.Text != "" {
		t.Errorf(`memo1.text isn't "", instead %v`, memo1.Text)
	}
}

func TestStoreMemo(t *testing.T) {
	var config config.Config
	workingDir, err := os.Getwd()
	if err != nil {
		t.Error("Error while query current working directory")
		return
	}
	config.TargetDir = workingDir + "/../../tmp"
	testFile := config.TargetDir + "/add_impl_test.md"
	// check if there is an old test file in the system
	_ = utils.DeleteFileIfExist(testFile)

	var memo Memo
	memo.Text = "I am from TestStoreMemo"
	memo.Targets = append(memo.Targets, "add_impl_test")
	err = StoreMemo(memo, config)
	if err != nil {
		t.Errorf("Error while store memo: %v", err)
	}
	b, err := utils.DoesFileExist(testFile)
	if (!b) || (err != nil) {
		t.Errorf("Seems that the memo file wasn't created: b=%v, err=%v", b, err)
	}
	if utils.DeleteFileIfExist(testFile) != nil {
		t.Errorf("Error while deleting tmp file output of the test: %s", testFile)
	}
}
