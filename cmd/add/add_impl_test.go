package add

import (
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
	if e != nil {
		t.Errorf("TestPartInput_2 error: %v", e)
	}
	l := len(ret.targets)
	if l != 1 {
		t.Errorf("TestPartInput_1 wrong target len: %d", l)
	}
	target := ret.targets[0]
	if target != "test" {
		t.Errorf("TestPartInput_1 wrong target content: %s", target)
	}

	text := ret.text
	if text != "aaa bbb" {
		t.Errorf("TestParamInput_1 wrong text: %s", text)
	}
}

func TestParseInput_2(t *testing.T) {
	var testData = []string{"aaa", "bbb"}
	ret, e := ParseInput(&testData)
	if e != nil {
		t.Errorf("TestPartInput_2 error: %v", e)
	}
	l := len(ret.targets)
	if l != 0 {
		t.Errorf("TestPartInput_2 wrong target len: %d", l)
	}
	text := ret.text
	if text != "aaa bbb" {
		t.Errorf("TestParamInput_2 wrong text: %s", text)
	}
}

func TestParseInput_3(t *testing.T) {
	var testData = []string{"#aaa", "#bbb"}
	ret, e := ParseInput(&testData)
	if e == nil {
		t.Errorf("TestPartInput_3 no error: %v", ret)
	}
	l := len(ret.targets)
	if l != 2 {
		t.Errorf("TestPartInput_3 wrong target len: %d", l)
	}
	text := ret.text
	if text != "" {
		t.Errorf("TestParamInput_3 wrong text: %s", text)
	}
}

func TestMemo(t *testing.T) {

	memo1 := Memo{}
	if memo1.text != "" {
		t.Errorf(`memo1.text isn't "", instead %v`, memo1.text)
	}
	if memo1.text != "" {
		t.Errorf(`memo1.text isn't "", instead %v`, memo1.text)
	}
}
