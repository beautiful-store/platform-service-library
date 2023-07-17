package apicall

import (
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeMessage(message interface{}) *APICall {
	var c APICall

	s, ok := message.(string)
	if !ok {
		fmt.Println("APICall DecodeMessage:", "not ok", "\n", message)
		return nil
	}
	err := lib.String2Struct(s, &c)
	if err != nil {
		fmt.Println("APICall DecodeMessage:", err.Error(), "\n", message)
		return nil
	}

	return &c
}
