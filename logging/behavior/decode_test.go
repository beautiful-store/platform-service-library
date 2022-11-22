package behavior

import (
	"testing"
)

func TestDecodeMessage(t *testing.T) {
	l := logContext{ModuleName: "testModule",
		TimeUnixNano:      1665620378841470000,
		Timestamp:         "2022-10-13T09:19:38+09:00",
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
