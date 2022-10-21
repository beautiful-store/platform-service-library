package sns

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/beautiful-store/platform-service-library/logging"
	"github.com/go-xorm/xorm"
)

type notification struct {
	Type             string `json:"Type"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	MessageID        string `json:"MessageId"`
	Signature        string `json:"Signature"`
	SignatureVersion string `json:"SignatureVersion"`
	SigningCertURL   string `json:"SigningCertURL"`
	SubscribeURL     string `json:"SubscribeURL"`
	Timestamp        string `json:"Timestamp"`
	Token            string `json:"Token"`
}

// revive:disable:unexported-return
func NewNotification(req *http.Request) *notification {
	b, _ := ioutil.ReadAll(req.Body)

	fmt.Println("1====", string(b))
	fmt.Println("1end====")

	reg1 := `(\\+)[n|t]`
	reg2 := `\\+`

	re1 := regexp.MustCompile(reg1)
	re2 := regexp.MustCompile(reg2)
	str1 := re1.ReplaceAllString(string(b), "")
	body := re2.ReplaceAllString(str1, "")

	// s, err := strconv.Unquote(string(b))
	// if err != nil {
	// 	fmt.Println("Unquote error=", err.Error())
	// 	return nil
	// }

	fmt.Println("2====", body)
	fmt.Println("2end====")
	m := notification{}
	if err := lib.Byte2Struct([]byte(body), &m); err != nil {
		fmt.Println("Byte2Struct error=", err.Error())
		return nil
	} else if m.MessageID == "" {
		fmt.Println("messageID error=", "can't conver to notification")
		return nil
	}

	return &m
}

func (n *notification) AddDB(engine *xorm.Engine) error {
	if n == nil {
		return errors.New("aws sns notification message is nil")
	}

	if n.Message == "" {
		return errors.New("aws sns notification message is empty")
	}

	log := logging.DecodeLogMessage(n.Message)
	if log == nil {
		return errors.New("log message converting error")
	}

	if err := log.CheckAndMakeTable(engine); err != nil {
		return err
	}

	if err := log.InsertTable(engine); err != nil {
		return err
	}

	return nil
}