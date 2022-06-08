package add

type MEMO struct {
	text    string
	targets []string
}

func IsInteractiveMode(args []string) bool {
	if len(args) == 0 {
		return true
	}
	return false
}

func ParseInput(inputStrings []string) (MEMO, error) {
	return MEMO{}, nil
}
