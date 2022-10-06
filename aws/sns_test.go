package aws

import "testing"

func TestIsEmailFormat(t *testing.T) {
	topic := "topic"
	NewSNS().WithTopic(topic).Send("this is test")
}
