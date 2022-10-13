package logging

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestRepository(t *testing.T) {
	origin := "@tcp(127.0.0.1:3306)/sharing?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"

	user := os.Getenv("LOCAL_DB_USER")
	pwd := os.Getenv("LOCAL_DB_PASSWORD")

	dbConnection := fmt.Sprintf("%s:%s%s", user, pwd, origin)

	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("Database open error: error: %s \n", err))
	}

	message := `{"module_name":"testModule","time_unix_nano":1665620378841470000,"timestamp":"2022-10-13T09:19:38+09:00","service_id":"","service_name":"UNKNOWN","parent_service_id":"","parent_service_name":"","remote_ip":"192.0.2.1","uri":"/","host":"example.com","method":"GET","path":"/","referer":"","user_agent":"","bytes_in":0,"bytes_out":0,"header":"","query":"","form":"","status":0,"panic":false,"error":"panic message","body":"","stack_trace":"","latency":0,"member_id":0,"member_orgid":0,"member_name":""}`
	log := DecodeLogMessage(message)

	if err = log.CheckAndMakeTable(engine); err != nil {
		t.Error(err)
	}

	if err = log.InsertTable(engine); err != nil {
		t.Error(err)
	}
}
