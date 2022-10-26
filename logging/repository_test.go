package logging

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestRepository(t *testing.T) {
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))
	json := `{"env":"dev","trace_id":"", "module_name":"testModule","time_unix_nano":1665713112206911000,"timestamp":"2022-10-14T11:05:12+09:00","service_id":"","service_name":"UNKNOWN","parent_service_id":"parentServiceID","parent_service_name":"parentServiceName","remote_ip":"192.0.2.1","uri":"/","host":"example.com","method":"GET","path":"/","referer":"","user_agent":"","bytes_in":0,"bytes_out":0,"header":"","query":"","form":"","status":0,"panic":false,"error":"","body":"","stack_trace":"","latency":0,"member_id":1,"member_orgid":1,"member_name":"member1"}`

	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	log := DecodeLogMessage(json)

	if err = log.CheckAndMakeTable(engine); err != nil {
		t.Error(err)
	}

	if err = log.InsertTable(engine); err != nil {
		t.Error(err)
	}
}
