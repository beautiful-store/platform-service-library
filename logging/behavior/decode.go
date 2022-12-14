package behavior

import (
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeMessage(message interface{}) *Log {
	c := logContext{}

	s, ok := message.(string)
	if ok {
		err := lib.String2Struct(s, &c)
		if err != nil {
			fmt.Println("Behavior DecodeMessage:", err.Error(), "\n", message)
		}
	}

	return &Log{Context: &c}
}
