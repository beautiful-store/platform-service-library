package login

import (
	"fmt"

	lib "github.com/beautiful-store/platform-service-library"
)

func DecodeMessage(message interface{}) *Login {
	var c Login

	s, ok := message.(string)
	if ok {
		err := lib.String2Struct(s, &c)
		if err != nil {
			fmt.Println("Login DecodeMessage:", err.Error(), "\n", message)
		}
	}

	return &c
}
