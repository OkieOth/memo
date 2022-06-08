package add

import "testing"

func TestParseInput(t *testing.T) {
	memo1 := MEMO{}
	if memo1.text != "" {
		t.Errorf(`memo1.text isn't "", instead %v`, memo1.text)
	}
}
