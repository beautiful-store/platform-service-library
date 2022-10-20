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

func TestNewNotificationAddDB(t *testing.T) {
	body := `"{\n \"Type\" : \"Notification\",\n \"MessageId\" : \"00000-0000-0000-0000-a0a000a00aa0\",\n \"TopicArn\" : \"arn:aws:sns:region:00000000:test_service\",\n \"Message\" : "{\\"env\\":\\"\\",\\"module_name\\":\\"testModule\\",\\"time_unix_nano\\":1666249515957513000,\\"timestamp\\":\\"2022-10-20T16:05:15+09:00\\",\\"service_id\\":\\"\\",\\"service_name\\":\\"UNKNOWN\\",\\"parent_service_id\\":\\"\\",\\"parent_service_name\\":\\"\\",\\"remote_ip\\":\\"192.0.2.1\\",\\"uri\\":\\"/\\",\\"host\\":\\"example.com\\",\\"method\\":\\"GET\\",\\"path\\":\\"/\\",\\"referer\\":\\"\\",\\"user_agent\\":\\"\\",\\"bytes_in\\":0,\\"bytes_out\\":0,\\"header\\":\\"\\",\\"query\\":\\"\\",\\"body\\":\\"\\",\\"status\\":0,\\"panic\\":false,\\"error\\":\\"\\",\\"stack_trace\\":\\"\\",\\"latency\\":0,\\"member_id\\":0,\\"member_orgid\\":0,\\"member_name\\":\\"\\"}",\n \"Timestamp\" : \"2022-10-20T05:25:54.442Z\",\n \"SignatureVersion\" : \"1\",\n \"Signature\" : \"uniqueno",\n \"SigningCertURL\" : \"https://sns.region.amazonaws.com/SimpleNotificationService-no.pem\",\n \"UnsubscribeURL\" : \"https://region.amazonaws.com/?Action=Unsubscribe\u0026SubscriptionArn=arn:aws:sns:region:000000000:test_service:00000-0000-0000-0000-a0a000a00aa0\"\n}"`

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
