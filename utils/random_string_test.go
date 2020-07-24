package utils

import "testing"

func TestRandomString(t *testing.T) {
	str, err := GenerateRandomString(4)
	if err != nil {
		t.Error(err)
	}
	if len(str) != 4 {
		t.Errorf("String is not 4 runes long! %+v", err)
	}
}
