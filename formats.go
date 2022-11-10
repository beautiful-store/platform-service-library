package library

import (
	"regexp"
	"time"
)

const (
	DateLayout8  = "20060102"
	DateLayout10 = "2006-01-02"
	DateLayout19 = "2006-01-02 15:04:05"
)

func GetTimeLocation(region string) *time.Location {
	loc, _ := time.LoadLocation(region)

	return loc
}

func GetAsiaSeoulTimeLocation() *time.Location {
	return GetTimeLocation("Asia/Seoul")
}

func IsEmailFormat(email string) bool {
	format := `^[a-zA-Z0-9+-\_.]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	matched, _ := regexp.MatchString(format, email)

	return matched
}
