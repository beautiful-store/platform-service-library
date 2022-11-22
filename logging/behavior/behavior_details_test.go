package behavior

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-xorm/xorm"
)

func TestConvertLogDetails(t *testing.T) {
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))

	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	logid := int64(999)
	str := `{"file":"/go/src/sharing-platform-service/donation/repository/donation_repository.go:270","func":"sharing-platform-service/donation/repository.donationRepository.FindAll","level":"trace","msg":"","time":"2022-10-21T15:07:43+09:00"}`
	stacks := ConvertLogDetails(logid, str)
	if len(stacks) == 0 {
		t.Error("converting error")
	} else {
		stacks.InsertTable(engine)
	}
}

func TestConvertLogSQLDetails(t *testing.T) {
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))

	engine, err := xorm.NewEngine("mysql", dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	logid := int64(999)
	str := `[info]  2022/11/22 14:26:40.612230 session_raw.go:43: [SQL] SELECT id, org_id, goods_id, member_id, delivery_name, delivery_mobile, delivery_post_no, delivery_address, delivery_address_detail, quantity, application_quantity, contents, target, upload_files, status, confirm_comments, confirmed, created, updated, application_date, delivered_date, del FROM applications WHERE (del is null or del = 0) AND (application_date between ? and ?) ORDER BY id DESC []interface {}{"20220630", "20220702"}`
	sqls := ConvertLogSQLDetails(logid, str)
	if len(sqls) == 0 {
		t.Error("converting error")
	} else {
		sqls.InsertTable(engine)
	}
}
