package library

import (
	"testing"
)

func TestContains(t *testing.T) {
	v := []string{"a", "b", "c"}
	v1 := "d"
	v2 := "c"

	boolean1 := Contains(v, v1)
	if boolean1 {
		t.Fatal(v, v1)
	}
	boolean2 := Contains(v, v2)
	if !boolean2 {
		t.Fatal(v, v2)
	}
}
