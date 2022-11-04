package sns

import (
	"fmt"
)

type sendTypeEnum int // API 유형

const (
	Behavior sendTypeEnum = 0 // 내부 api 콜
	APICall  sendTypeEnum = 1 // 내부 api 콜
)

var sendTypeEnumValues = [...]string{"Behavior", "APICall"}

func (s sendTypeEnum) String() string {
	switch {
	case int(s) <= 1:
		return sendTypeEnumValues[s]
	default:
		return fmt.Sprintf("%d", s)
	}
}
