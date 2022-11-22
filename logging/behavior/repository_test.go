package behavior

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestRepository(t *testing.T) {
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))
	// json := `{"env":"dev","trace_id":"", "module_name":"testModule","time_unix_nano":1665713112206911000,"timestamp":"2022-10-14T11:05:12+09:00","service_id":"","service_name":"UNKNOWN","parent_service_id":"parentServiceID","parent_service_name":"parentServiceName","remote_ip":"192.0.2.1","uri":"/","host":"example.com","method":"GET","path":"/","referer":"","user_agent":"","bytes_in":0,"bytes_out":0,"header":"","query":"","form":"","status":0,"panic":false,"error":"","body":"","stack_trace":"","latency":0,"member_id":1,"member_orgid":1,"member_name":"member1"}`
	l := logContext{ModuleName: "testModule",
		TimeUnixNano:      1665620378841470000,
		Timestamp:         "2022-10-13T09:19:38+09:00",
		ServiceID:         "",
		ServiceName:       "UNKNOWN",
		ParentServiceID:   "",
		ParentServiceName: "",
		RemoteIP:          "192.0.2.1",
		URI:               "/",
		Host:              "example.com",
		Method:            "GET",
		Path:              "/",
		Referer:           "",
		UserAgent:         "",
		BytesIn:           0,
		BytesOut:          0,
		Header:            "",
		Query:             "",
		Status:            0,
		Panic:             false,
		Error:             "panic message",
		Body:              "",
		StackTrace:        "",
		SQLTrace:          "",
		Latency:           0,
		MemberID:          0,
		MemberOrgID:       0,
		MemberName:        "",
	}

	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	log := DecodeMessage(l)

	log.CheckAndMakeTable(engine)
	if err = log.InsertTable(engine); err != nil {
		t.Error(err)
	}
}
