package behavior

import (
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	l := logContext{ModuleName: "testModule",
		TimeUnixNano:      1665620378841470000,
		Timestamp:         "2006-01-02T15:04:05.999999999Z07:00",
		ServiceID:         "",
		ServiceName:       "UNKNOWN",
		ParentServiceID:   "",
		ParentServiceName: "",
		RemoteIP:          "192.0.2.1",
		URI:               "/",
		Host:              "example.com",
		Method:            "GET",
		Path:              "/",
		Referer:           "",
		UserAgent:         "",
		BytesIn:           0,
		BytesOut:          0,
		Header:            "",
		Query:             "",
		Status:            0,
		Panic:             false,
		Error:             "panic message",
		Body:              "",
		StackTrace:        "",
		SQLTrace:          "",
		Latency:           0,
		MemberID:          0,
		MemberOrgID:       0,
		MemberName:        "",
	}

	log := DecodeMessage(l)

	log.OutToConsole()
	t.Log(log.Context)
}
