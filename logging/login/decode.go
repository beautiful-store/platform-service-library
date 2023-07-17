package login

import (
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeMessage(message interface{}) *Login {
	var c Login

	s, ok := message.(string)
	if !ok {
		fmt.Println("Login DecodeMessage:", "not ok", "\n", message)
		return nil
	}
	err := lib.String2Struct(s, &c)
	if err != nil {
		fmt.Println("Login DecodeMessage:", err.Error(), "\n", message)
		return nil
	}

	return &c
}
