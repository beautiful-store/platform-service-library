package logging

import (
	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeLogMessage(message string) *Log {
	c := logContext{}
	if err := lib.String2Struct(message, &c); err != nil {
		panic(err.Error())
	}

	return &Log{Context: &c}
}
