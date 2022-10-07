package aws

import (
	"testing"
)

func TestIsEmailFormat(t *testing.T) {
	topic := "topic"
	_, err := NewSNS().WithTopic(topic).Send("this is test")
	if err != nil {
		t.Fatal(err.Error())
	}
}
