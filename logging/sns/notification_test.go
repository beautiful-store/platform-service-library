package sns

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestNewNotificationAddDB_Behavior(t *testing.T) {
	body := `{
		"Type" : "Behavior",
		"MessageId" : "0000000-0000-0000-0000-000000000000",
		"TopicArn" : "arn:aws:sns:region:000000000000:topic",
		"Message" : "{\"env\":\"dev\",\"module_name\":\"service\",\"time_unix_nano\":1666328081182871206,\"timestamp\":\"2022-10-21T13:54:41+09:00\",\"service_id\":\"unique_no\",\"service_name\":\"service_name\",\"parent_service_id\":\"\",\"parent_service_name\":\"\",\"remote_ip\":\"0.0.0.0\",\"uri\":\"/api\",\"host\":\"host\",\"method\":\"GET\",\"path\":\"/api\",\"referer\":\"\",\"user_agent\":\"PostmanRuntime/7.29.2\",\"bytes_in\":1,\"bytes_out\":1,\"header\":\"\",\"query\":\"\",\"body\":\"{\\\"}\",\"status\":200,\"panic\":false,\"error\":\"\",\"stack_trace\":\"\",\"latency\":1,\"member_id\":0,\"member_orgid\":0,\"member_name\":\"\"}",
		"Timestamp" : "2022-10-21T04:54:41.207Z",
		"SignatureVersion" : "1",
		"Signature" : "signature",
		"SigningCertURL" : "https://sns.region/SimpleNotificationService",
		"UnsubscribeURL" : "https://sns.region/?"
	}`

	reqBody := bytes.NewBufferString(body)
	req := httptest.NewRequest("post", "https://dev-share-service.beautiful0.org/api/logging", reqBody)

	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOG_DB_CONNECTION"))
	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	n := NewNotification(req)
	err = n.AddDB(engine)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNewNotificationAddDB_BehaviorDetail(t *testing.T) {
	body := `{
		"Type" : "BehaviorDetail",
		"MessageId" : "0000000-0000-0000-0000-000000000000",
		"TopicArn" : "arn:aws:sns:region:000000000000:topic",
		"Message" : "{\"file\":\"/Users/auth.go:01\",\"func\":\"member-platform-service/handlers.Auth.GetLoginInfo\",\"level\":\"trace\",\"msg\":\"\",\"service_id\":\"gt4s2Nabya36kF0aFBO8ZCs2y07RaXLp\",\"time\":\"2023-06-16T10:39:20.575527+09:00\"}",
		"Timestamp" : "2022-10-21T04:54:41.207Z",
		"SignatureVersion" : "1",
		"Signature" : "signature",
		"SigningCertURL" : "https://sns.region/SimpleNotificationService",
		"UnsubscribeURL" : "https://sns.region/?"
	}`

	reqBody := bytes.NewBufferString(body)
	req := httptest.NewRequest("post", "https://dev-share-service.beautiful0.org/api/logging", reqBody)

	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))
	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	n := NewNotification(req)
	err = n.AddDB(engine)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestNewNotificationAddDB_Apicall(t *testing.T) {
	body := `{
		"Type" : "ApiCall",
		"MessageId" : "0000000-0000-0000-0000-000000000000",
		"TopicArn" : "arn:aws:sns:region:000000000000:topic",
		"Message" : "{\"env\":\"dev\",\"module_name\":\"service\",\"time_unix_nano\":1666328081182871206,\"timestamp\":\"2022-10-21T13:54:41+09:00\",\"service_id\":\"unique_no\",\"service_name\":\"service_name\",\"parent_service_id\":\"\",\"parent_service_name\":\"\",\"remote_ip\":\"0.0.0.0\",\"uri\":\"/api\",\"host\":\"host\",\"method\":\"GET\",\"path\":\"/api\",\"referer\":\"\",\"user_agent\":\"PostmanRuntime/7.29.2\",\"bytes_in\":1,\"bytes_out\":1,\"header\":\"\",\"query\":\"\",\"body\":\"{\\\"}\",\"status\":200,\"panic\":false,\"error\":\"\",\"stack_trace\":\"\",\"latency\":1,\"member_id\":0,\"member_orgid\":0,\"member_name\":\"\"}",
		"Timestamp" : "2022-10-21T04:54:41.207Z",
		"SignatureVersion" : "1",
		"Signature" : "signature",
		"SigningCertURL" : "https://sns.region/SimpleNotificationService",
		"UnsubscribeURL" : "https://sns.region/?"
	}`

	reqBody := bytes.NewBufferString(body)
	req := httptest.NewRequest("post", "https://dev-share-service.beautiful0.org/api/logging", reqBody)

	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))
	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	n := NewNotification(req)
	err = n.AddDB(engine)
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestNewNotificationAddDB_login(t *testing.T) {
	body := `{
		"Type" : "LogIn",
		"MessageId" : "0000000-0000-0000-0000-000000000000",
		"TopicArn" : "arn:aws:sns:region:000000000000:topic",
		"Message" : "{\"env\":\"dev\",\"module_name\":\"service\",\"time_unix_nano\":1666328081182871206,\"timestamp\":\"2022-10-21T13:54:41+09:00\",\"service_id\":\"unique_no\",\"service_name\":\"service_name\",\"parent_service_id\":\"\",\"parent_service_name\":\"\",\"remote_ip\":\"0.0.0.0\",\"uri\":\"/api\",\"host\":\"host\",\"method\":\"GET\",\"path\":\"/api\",\"referer\":\"\",\"user_agent\":\"PostmanRuntime/7.29.2\",\"bytes_in\":1,\"bytes_out\":1,\"header\":\"\",\"query\":\"\",\"body\":\"{\\\"}\",\"status\":200,\"panic\":false,\"error\":\"\",\"stack_trace\":\"\",\"latency\":1,\"member_id\":0,\"member_orgid\":0,\"member_name\":\"\"}",
		"Timestamp" : "2022-10-21T04:54:41.207Z",
		"SignatureVersion" : "1",
		"Signature" : "signature",
		"SigningCertURL" : "https://sns.region/SimpleNotificationService",
		"UnsubscribeURL" : "https://sns.region/?"
	}`

	reqBody := bytes.NewBufferString(body)
	req := httptest.NewRequest("post", "https://dev-share-service.beautiful0.org/api/logging", reqBody)

	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))
	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	n := NewNotification(req)
	err = n.AddDB(engine)
	if err != nil {
		t.Fatal(err.Error())
	}
}
