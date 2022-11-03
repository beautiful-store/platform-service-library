package apicall

import (
	"fmt"
)

type typeEnum int // API 유형

const (
	InternalShaingAPI     typeEnum = 0 // 내부 api 콜
	InternalOrgSharingAPI typeEnum = 1 // 내부 api 콜
	InternalTaggingAPI    typeEnum = 2 // 내부 api 콜
	KakaoLogin            typeEnum = 3 // kakao 로그인
	KakaoBusinessTalk     typeEnum = 4 // kakao 알람톡
	KakaoFriendTalk       typeEnum = 5 // kakao 친구톡
	ExternalCourierCU     typeEnum = 6 // 택배 CU API
	ExternalCourierGS     typeEnum = 7 // 택배 GS API
)

var apiCallTypeEnumValues = [...]string{
	"InternalShaingAPI", "InternalOrgSharingAPI", "InternalTaggingAPI", // 내부프로젝트
	"KakaoLogin", "KakaoBusinessTalk", "KakaoFriendTalk", // 카카오 관련
	"ExternalCourierCU", "ExternalCourierGS", // 택배 관련
}

func (s typeEnum) String() string {
	switch {
	case int(s) <= 7:
		return apiCallTypeEnumValues[s]
	default:
		return fmt.Sprintf("%d", s)
	}
}
