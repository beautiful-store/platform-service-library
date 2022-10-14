package aws

import (
	"os"
	"testing"
)

func TestSendSNS(t *testing.T) {
	topic := os.Getenv("AWSSNS_BEHAVIOR_LOG_TOPIC")
	if topic == "" {
		topic = "arn:aws:sns:region:123456789:test_topic"
	}
	_, err := NewSNS().WithTopic(topic).Send("this is test")
	if err != nil {
		t.Fatal(err.Error())
	}
}
