package add

import (
	"bytes"
	"testing"
)

func TestIsInteractiveMode(t *testing.T) {
	empty := make([]string, 0)
	if !IsInteractiveMode(&empty) {
		t.Error("Interactive mode isn't triggered by empty array")
	}

	oneField := make([]string, 1)
	oneField[0] = "test"
	if IsInteractiveMode(&oneField) {
		t.Error("Interactive mode triggered by non empty array")
	}
}

func TestParseInput_1(t *testing.T) {
	var testData = []string{"aaa", "bbb", "#test"}
	ret, e := ParseInput(&testData)
	parseInput1(ret, e, t, "TestParseInput_1")
}

func TestGetMemoFromStdin_1(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("aaa bbb #test\n"))
	ret, e := GetMemoFromStdin(&stdin)
	parseInput1(ret, e, t, "TestGetMemoFromStdin_1")
}

func TestParseInput_2(t *testing.T) {
	var testData = []string{"aaa", "bbb"}
	ret, e := ParseInput(&testData)
	parseInput2(ret, e, t, "TestParseInput_2")
}

func TestGetMemoFromStdin_2(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("aaa bbb\n"))
	ret, e := GetMemoFromStdin(&stdin)
	parseInput2(ret, e, t, "TestGetMemoFromStdin_2")
}

func TestParseInput_3(t *testing.T) {
	var testData = []string{"#aaa", "#bbb"}
	ret, e := ParseInput(&testData)
	parseInput3(ret, e, t, "TestParseInput_3")
}

func TestGetMemoFromStdin_3(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("#aaa #bbb\n"))
	ret, e := GetMemoFromStdin(&stdin)
	parseInput3(ret, e, t, "TestGetMemoFromStdin_3")
}

func TestParseInput_4(t *testing.T) {
	var testData = []string{"aaa", "bbb", "#test", "`xxx`,**bold**,("}
	ret, e := ParseInput(&testData)
	parseInput4(ret, e, t, "TestParseInput_4")
}

func TestGetMemoFromStdin_4(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("aaa bbb #test `xxx`,**bold**,(\n"))
	ret, e := GetMemoFromStdin(&stdin)
	parseInput4(ret, e, t, "TestGetMemoFromStdin_4")
}

func parseInput1(ret Memo, e error, t *testing.T, caller string) {
	expectedTargets := []string{"test"}
	parseInput(ret, e, t, caller, "aaa bbb", expectedTargets, false)
}

func parseInput2(ret Memo, e error, t *testing.T, caller string) {
	expectedTargets := []string{}
	parseInput(ret, e, t, caller, "aaa bbb", expectedTargets, false)
}

func parseInput3(ret Memo, e error, t *testing.T, caller string) {
	expectedTargets := []string{"aaa", "bbb"}
	parseInput(ret, e, t, caller, "", expectedTargets, true)
}

func parseInput4(ret Memo, e error, t *testing.T, caller string) {
	expectedTargets := []string{"test"}
	parseInput(ret, e, t, caller, "aaa bbb `xxx`,**bold**,(", expectedTargets, false)
}

func parseInput(ret Memo, e error, t *testing.T, caller string, expectedTxt string, expectedTargets []string, errorExpected bool) {
	if (!errorExpected) && (e != nil) {
		t.Errorf("%s error: %v", caller, e)
	}
	l := len(ret.Targets)
	expectedTargetsLen := len(expectedTargets)
	if l != expectedTargetsLen {
		t.Errorf("%s wrong target len: %d, expected len: %d", caller, l, expectedTargetsLen)
	}
	for i, expectedTarget := range expectedTargets {
		if expectedTarget != ret.Targets[i] {
			t.Errorf("%s wrong target: %s, expected: %s", caller, ret.Targets[i], expectedTarget)
		}
	}
	text := ret.Text
	if text != expectedTxt {
		t.Errorf("%s wrong text: %s, expected: %s", caller, text, expectedTxt)
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
