package logging

import (
	"testing"

	lib "github.com/beautiful-store/platform-service-library"
)

func TestDecodeLogMessage(t *testing.T) {
	json := `{"module_name":"testModule","time_unix_nano":1665620378841470000,"timestamp":"2022-10-13T09:19:38+09:00","service_id":"","service_name":"UNKNOWN","parent_service_id":"","parent_service_name":"","remote_ip":"192.0.2.1","uri":"/","host":"example.com","method":"GET","path":"/","referer":"","user_agent":"","bytes_in":0,"bytes_out":0,"header":"","query":"","form":"","status":0,"panic":false,"error":"panic message","body":"","stack_trace":"","latency":0,"member_id":0,"member_orgid":0,"member_name":""}`

	log := DecodeLogMessage(json)

	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}
