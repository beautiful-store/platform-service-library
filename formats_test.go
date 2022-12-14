package library

import "testing"

func TestIsEmailFormat(t *testing.T) {
	v1 := []string{"a.bcde", "abcde@", "@bcde", "aaaaaa@bbbbb"}

	for _, email := range v1 {
		boolean := IsEmailFormat(email)
		if boolean {
			t.Fatal(email, " check error")
		}
	}
}

func TestGetDefaultLogLocalDateTimeMilli(t *testing.T) {
	if len(GetDefaultLogLocalDateTimeMilli()) != 29 {
		t.Error(GetDefaultLogLocalDateTimeMilli())
	}
}
