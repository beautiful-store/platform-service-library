package logging

import (
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeLogMessage(message string) *Log {
	c := logContext{}
	if err := lib.String2Struct(message, &c); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &Log{Context: &c}
}
