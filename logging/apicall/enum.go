package apicall

import (
	"fmt"
)

type typeEnum int // API 유형

const (
	InternalShaingAPI     typeEnum = 0 // 내부 api 콜
	InternalOrgSharingAPI typeEnum = 1 // 내부 api 콜
	InternalTaggingAPI    typeEnum = 2 // 내부 api 콜
	InternalMemberAPI     typeEnum = 3 // 내부 api 콜
	InternalAccountAPI    typeEnum = 4 // 내부 api 콜
	KakaoLogin            typeEnum = 5 // kakao 로그인
	KakaoBusinessTalk     typeEnum = 6 // kakao 알람톡
	KakaoFriendTalk       typeEnum = 7 // kakao 친구톡
	ExternalCourierCU     typeEnum = 8 // 택배 CU API
	ExternalCourierGS     typeEnum = 9 // 택배 GS API

)

var apiCallTypeEnumValues = [...]string{
	"InternalShaingAPI", "InternalOrgSharingAPI", "InternalTaggingAPI", "InternalMemberAPI", "InternalAccountAPI", // 내부프로젝트
	"KakaoLogin", "KakaoBusinessTalk", "KakaoFriendTalk", // 카카오 관련
	"ExternalCourierCU", "ExternalCourierGS", // 택배 관련
}

func (s typeEnum) String() string {
	switch {
	case int(s) <= 9:
		return apiCallTypeEnumValues[s]
	default:
		return fmt.Sprintf("%d", s)
	}
}
