package behavior

import (
	"encoding/json"
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeMessage(message interface{}) *Log {
	c := logContext{}

	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Behavior DecodeMessage:", "not ok", "\n", err, message)
		return nil
	}
	// err := lib.String2Struct(s, &c)
	if err := lib.Byte2Struct(jsonData, &c); err != nil {
		fmt.Println("Behavior DecodeMessage:", err.Error(), "\n", message)
		return nil
	}

	return &Log{Context: &c}
}

func DecodeDetailMessage(message interface{}, l interface{}) error {
	s, ok := message.(string)
	if !ok {
		err := lib.String2Struct(s, l)
		if err != nil {
			fmt.Println(" DecodeMessage:", err.Error(), "\n", message)
		}
	}

	return nil
}
