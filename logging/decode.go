package logging

import (
	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeLogMessage(message string) *Log {
	c := logContext{}
	lib.String2Struct(message, &c)

	return &Log{Context: &c}
}
