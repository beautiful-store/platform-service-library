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
	json := `{"env":"local","module_name":"sharing-platform-service","time_unix_nano":1666165940791718000,"timestamp":"2022-10-19T16:52:20+09:00","service_id":"iHnjjKBVZf0IdMi0kb8UkGw3BfItK1vw","service_name":"sharing-platform-service/controller.DonationAdminController.CreateWithClassificationByVisit-fm","parent_service_id":"","parent_service_name":"","remote_ip":"::1","uri":"/api/admin/donations/visit/classifications","host":"localhost:7000","method":"POST","path":"/api/admin/donations/visit/classifications","referer":"","user_agent":"PostmanRuntime/7.29.2","bytes_in":281,"bytes_out":3346,"header":"","query":"","form":"","status":200,"panic":false,"error":"","body":""
  ,"stack_trace":"{\"file\":\"/Users/chaos/go/sharing-platform-service/controller/donation_admin_controller.go:76\",\"func\":\"sharing-platform-service/controller.DonationAdminController.CreateWithClassificationByVisit\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/service/donation_validation_service.go:20\",\"func\":\"sharing-platform-service/donation/service.donationService.ValidateCampaign\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/service/donation_validation_service.go:107\",\"func\":\"sharing-platform-service/donation/service.donationService.ValidateSite\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/service/donation_service.go:85\",\"func\":\"sharing-platform-service/donation/service.donationService.CreateWithClassification\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_entity.go:103\",\"func\":\"sharing-platform-service/donation/entity.(*Donation).CreateWithClassification\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_entity.go:54\",\"func\":\"sharing-platform-service/donation/entity.(*Donation).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_entity.go:57\",\"func\":\"sharing-platform-service/donation/entity.(*Donation).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donations (member_id,campaign_id,campaign_name,status,donation_date,name,mobile,site_type,site_code,site_Name,take_over_method_type,agreed,marketing_agreed,sharing_type,eval_type,appended,created,updated,confirmed) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)[1 4 아름다운가게 매장방문 기부 DonationRegistered 20221019 aaa DxLjQyB85DTr0rU= STORE 200020 양재점 VISIT true false 100 AMOUNT false {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} null]\",\"sql\":true,\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_status_revision_entity.go:26\",\"func\":\"sharing-platform-service/donation/entity.(*DonationStatusRevision).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_status_revision_entity.go:29\",\"func\":\"sharing-platform-service/donation/entity.(*DonationStatusRevision).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donation_status_revisions (donation_id,status,created) VALUES (?, ?, ?)[319175 DonationRegistered {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"}]\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:70\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:73\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donation_classification_details (donation_id,item_type,unit_price,quantity,note,created,updated) VALUES (?, ?, ?, ?, ?, ?, ?)[319175 100 1318 1  {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"}]\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:70\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:73\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donation_classification_details (donation_id,item_type,unit_price,quantity,note,created,updated) VALUES (?, ?, ?, ?, ?, ?, ?)[319175 200 1758 2  {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"}]\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:70\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:73\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donation_classification_details (donation_id,item_type,unit_price,quantity,note,created,updated) VALUES (?, ?, ?, ?, ?, ?, ?)[319175 300 491 3  {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"}]\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:70\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/donation/entity/donation_classification_detail_entity.go:73\",\"func\":\"sharing-platform-service/donation/entity.(*DonationClassificationDetail).Create\",\"level\":\"info\",\"msg\":\"INSERT INTO donation_classification_details (donation_id,item_type,unit_price,quantity,note,created,updated) VALUES (?, ?, ?, ?, ?, ?, ?)[319175 400 5756 4  {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"} {\\\"id\\\":19,\\\"role\\\":\\\"시스템 관리자,조직/멤버 관리자,모바일사업본부\\\",\\\"orgid\\\":0,\\\"name\\\":\\\"유선희\\\",\\\"datetime\\\":\\\"2022-10-19 16:52:20\\\"}]\",\"time\":\"2022-10-19T16:52:20+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/service/http_interface_log_service.go:33\",\"func\":\"sharing-platform-service/adapter/service.(*httpInterfaceLogService).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/repository/http_interface_log_repository.go:30\",\"func\":\"sharing-platform-service/adapter/repository.httpInterfaceLog.Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/repository/http_interface_log_repository.go:36\",\"func\":\"sharing-platform-service/adapter/repository.httpInterfaceLog.Create\",\"level\":\"info\",\"msg\":\"1 was affected (http_interface_logs)\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/kakao/adapter/kakao_bizmessage_adapter.go:218\",\"func\":\"sharing-platform-service/kakao/adapter.kakaoBizmessageAdapter.SendDonationRegisteredToManagerMessage\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/service/http_interface_log_service.go:33\",\"func\":\"sharing-platform-service/adapter/service.(*httpInterfaceLogService).Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/repository/http_interface_log_repository.go:30\",\"func\":\"sharing-platform-service/adapter/repository.httpInterfaceLog.Create\",\"level\":\"trace\",\"msg\":\"\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n,{\"file\":\"/Users/chaos/go/sharing-platform-service/adapter/repository/http_interface_log_repository.go:36\",\"func\":\"sharing-platform-service/adapter/repository.httpInterfaceLog.Create\",\"level\":\"info\",\"msg\":\"1 was affected (http_interface_logs)\",\"time\":\"2022-10-19T16:52:21+09:00\"}\n","latency":635,"member_id":19,"member_orgid":0,"member_name":""}`

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
