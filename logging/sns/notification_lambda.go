package sns

// import (
// 	"errors"

// 	"github.com/go-xorm/xorm"
// )

// type NotificationLambda struct {
// 	Type             string `json:"Type"`
// 	TopicArn         string `json:"TopicArn"`
// 	Message          string `json:"Message"`
// 	MessageID        string `json:"MessageId"`
// 	Signature        string `json:"Signature"`
// 	SignatureVersion string `json:"SignatureVersion"`
// 	SigningCertURL   string `json:"SigningCertURL"`
// 	SubscribeURL     string `json:"SubscribeURL"`
// 	Timestamp        string `json:"Timestamp"`
// 	Token            string `json:"Token"`
// }

// func (lambdaNoti *NotificationLambda) AddDB(engine *xorm.Engine) error {
// 	if lambdaNoti.Message == "" {
// 		return errors.New("aws sns notification message is empty")
// 	}

// 	n := notification{
// 		Type:             lambdaNoti.Type,
// 		TopicArn:         lambdaNoti.TopicArn,
// 		Message:          lambdaNoti.Message,
// 		MessageID:        lambdaNoti.MessageID,
// 		Signature:        lambdaNoti.Signature,
// 		SignatureVersion: lambdaNoti.SignatureVersion,
// 		SigningCertURL:   lambdaNoti.SigningCertURL,
// 		SubscribeURL:     lambdaNoti.SubscribeURL,
// 		Timestamp:        lambdaNoti.Timestamp,
// 		Token:            lambdaNoti.Token,
// 	}

// 	return n.AddDB(engine)
// }