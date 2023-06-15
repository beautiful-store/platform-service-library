package sns

import (
	"fmt"
)

type sendTypeEnum int // API 유형

const (
	Behavior       sendTypeEnum = 0 // api service access log
	APICall        sendTypeEnum = 1 // 내부 api 콜
	LogIn          sendTypeEnum = 2 // 사용자 로그인
	BehaviorDetail sendTypeEnum = 3 // trace log
	BehaviorSql    sendTypeEnum = 4 // sql log
)

var sendTypeEnumValues = [...]string{"Behavior", "APICall", "LogIn", "BehaviorDetail", "BehaviorSql"}

func (s sendTypeEnum) String() string {
	switch {
	case int(s) <= 4:
		return sendTypeEnumValues[s]
	default:
		return fmt.Sprintf("%d", s)
	}
}
