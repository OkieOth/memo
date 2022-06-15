package utils

import (
	"testing"
)

func TestReplaceEnvVars(t *testing.T) {
	ret, err := ReplaceEnvVars("$HOME/$USER/${PATH}")
	if err != nil {
		t.Errorf("Error in ReplaceEnvVar: %v", err)
		return
	}
	if ret == "$HOME/$USER/${UID}" {
		t.Errorf("Nothing replaced: %v", ret)
		return
	}
}

func TestExtractEnvVars(t *testing.T) {
	output := extractEnvVars("HOME/.memo/targets")
	outputLen := len(output)
	if outputLen != 0 {
		t.Errorf("Wrong number of env vars found (0): %d", outputLen)
		return
	}

	output1 := extractEnvVars("$HOME/.memo/targets")
	outputLen1 := len(output1)
	if outputLen1 != 1 {
		t.Errorf("Wrong number of env vars found (1): %d", outputLen1)
		return
	}
	expected := "$HOME"
	if output1[0] != expected {
		t.Errorf("output1[0] != %s, %s", output1[0], expected)
		return
	}

	output2 := extractEnvVars("${HOME}/.memo/targets$HOMER")
	outputLen2 := len(output2)
	if outputLen2 != 2 {
		t.Errorf("Wrong number of env vars found (2): %d", outputLen2)
		return
	}
	expected = "${HOME}"
	if output2[0] != expected {
		t.Errorf("output2[0] != %s, %s", output2[0], expected)
		return
	}
	expected = "$HOMER"
	if output2[1] != expected {
		t.Errorf("output2[1] != %s, %s", output2[1], expected)
		return
	}

	output3 := extractEnvVars("test${1}${HOME}/.memo/targets$HOMER")
	outputLen3 := len(output3)
	if outputLen3 != 2 {
		t.Errorf("Wrong number of env vars found (3): %d", outputLen3)
		return
	}
	expected = "${HOME}"
	if output3[0] != expected {
		t.Errorf("output3[0] != %s, %s", output3[0], expected)
		return
	}
	expected = "$HOMER"
	if output3[1] != expected {
		t.Errorf("output3[1] != %s, %s", output3[1], expected)
		return
	}

	output4 := extractEnvVars("test${X1}${HOME}/.memo/targets$HOMER")
	outputLen4 := len(output4)
	if outputLen4 != 3 {
		t.Errorf("Wrong number of env vars found (4): %d", outputLen4)
		return
	}
	expected = "${X1}"
	if output4[0] != expected {
		t.Errorf("output4[0] != %s, %s", output4[0], expected)
		return
	}
	expected = "${HOME}"
	if output4[1] != expected {
		t.Errorf("output4[1] != %s, %s", output4[1], expected)
		return
	}
	expected = "$HOMER"
	if output4[2] != expected {
		t.Errorf("output4[2] != %s, %s", output4[2], expected)
		return
	}
}
