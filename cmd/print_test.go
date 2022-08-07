package cmd

import (
	"bufio"
	"os"
	"testing"
)

func TestGetTargetShortName(t *testing.T) {
	s1 := getTargetShortName("/home/eiko/.memo/targets/default.md")
	if s1 != "default" {
		t.Errorf("Wrong Text ... expected 'default' got '%s'", s1)
	}
}

func TestGetLastFileEntries10(t *testing.T) {
	file, err := os.Open("../internal/resources/dummy_target_content.md")
	if err != nil {
		t.Errorf("ERROR while reading file with test data: %v", err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	last10Entries := getLastFileEntries(fileScanner, 10)
	len10 := len(*last10Entries)
	if len10 != 10 {
		t.Errorf("return array don't consist of 10 elems: %d", len10)
		return
	}
	if (*last10Entries)[0] != "* 144 -T xxx [202206236252531]" ||
		(*last10Entries)[1] != "* 145 -T xxx [202206236262625]" ||
		(*last10Entries)[2] != "* 146 -t ddd [202206236262646]" ||
		(*last10Entries)[3] != "* 147 -t aaa [202206236282816]" ||
		(*last10Entries)[4] != "* 148 --text=xxx [202206236282847]" ||
		(*last10Entries)[5] != "* 149 --text=xxx [202206236292906]" ||
		(*last10Entries)[6] != "* 150 --text=xxx [202206236323201]" ||
		(*last10Entries)[7] != "* 151 --text=xxx [202206236424248]" ||
		(*last10Entries)[8] != "* 152 -text xxx [202206236434355]" ||
		(*last10Entries)[9] != "* 153 Das ist ein Text [202206266313119]" {
		t.Errorf("wrong last 10 entries: %v", last10Entries)
		return
	}
}

func TestGetLastFileEntries155(t *testing.T) {
	file, err := os.Open("../internal/resources/dummy_target_content.md")
	if err != nil {
		t.Errorf("ERROR while reading file with test data: %v", err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	last155Entries := getLastFileEntries(fileScanner, 155)
	len155 := len(*last155Entries)
	if len155 != 155 {
		t.Errorf("return array don't consist of 155 elems: %d", len155)
		return
	}
	if (*last155Entries)[0] == "" ||
		(*last155Entries)[152] == "" ||
		(*last155Entries)[153] != "" ||
		(*last155Entries)[154] != "" {
		t.Errorf("wrong last 155 entries: %v", last155Entries)
		return
	}
}

func TestGetLastFileEntries10Small(t *testing.T) {
	file, err := os.Open("../internal/resources/dummy_target_content_small.md")
	if err != nil {
		t.Errorf("ERROR while reading file with test data: %v", err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	last10Entries := getLastFileEntries(fileScanner, 10)
	len10 := len(*last10Entries)
	if len10 != 10 {
		t.Errorf("return array don't consist of 10 elems: %d", len10)
		return
	}
	if (*last10Entries)[0] != "* 001 Das ist ein kleiner Test [202206226101035]" ||
		(*last10Entries)[1] != "* 002 Das ist ein kleiner Test [202206226101035]" ||
		(*last10Entries)[2] != "* 003 Das ist ein kleiner Test [202206226111113]" ||
		(*last10Entries)[3] != "" ||
		(*last10Entries)[4] != "" ||
		(*last10Entries)[5] != "" ||
		(*last10Entries)[6] != "" ||
		(*last10Entries)[7] != "" ||
		(*last10Entries)[8] != "" ||
		(*last10Entries)[9] != "" {
		t.Errorf("wrong last 10 entries: %v", last10Entries)
		return
	}
}

func TestGetLastFileEntries3(t *testing.T) {
	file, err := os.Open("../internal/resources/dummy_target_content.md")
	if err != nil {
		t.Errorf("ERROR while reading file with test data: %v", err)
		return
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	last3Entries := getLastFileEntries(fileScanner, 3)
	len3 := len(*last3Entries)
	if len3 != 3 {
		t.Errorf("return array doesn't consist of 3 elems: %d", len3)
		return
	}
	if (*last3Entries)[0] != "* 151 --text=xxx [202206236424248]" ||
		(*last3Entries)[1] != "* 152 -text xxx [202206236434355]" ||
		(*last3Entries)[2] != "* 153 Das ist ein Text [202206266313119]" {
		t.Errorf("wrong last 3 entries: %v", last3Entries)
		return
	}

}
