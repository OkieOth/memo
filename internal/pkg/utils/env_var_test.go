package utils

import (
	"testing"
)

func TestReplaceEnvVars(t *testing.T) {
	output := ReplaceEnvVars("HOME/.memo/targets")
	outputLen := len(output)
	if outputLen != 0 {
		t.Errorf("Wrong number of env vars found (0): %d", outputLen)
	}

	output1 := ReplaceEnvVars("$HOME/.memo/targets")
	outputLen1 := len(output1)
	if outputLen1 != 1 {
		t.Errorf("Wrong number of env vars found (1): %d", outputLen1)
	}

	output2 := ReplaceEnvVars("${HOME}/.memo/targets$HOMER")
	outputLen2 := len(output2)
	if outputLen2 != 2 {
		t.Errorf("Wrong number of env vars found (2): %d", outputLen2)
	}

	output3 := ReplaceEnvVars("test${1}${HOME}/.memo/targets$HOMER")
	outputLen3 := len(output3)
	if outputLen3 != 2 {
		t.Errorf("Wrong number of env vars found (3): %d", outputLen3)
	}

	output4 := ReplaceEnvVars("test${X1}${HOME}/.memo/targets$HOMER")
	outputLen4 := len(output4)
	if outputLen4 != 3 {
		t.Errorf("Wrong number of env vars found (4): %d", outputLen4)
	}
}
