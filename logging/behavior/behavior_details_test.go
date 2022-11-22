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
	str := `{"file":"/go/src/service/test/repository/test_repository.go:1","func":"service/test/repository.testRepository.FindAll","level":"trace","msg":"","time":"2022-10-21T15:07:43+09:00"}`
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
	str := `[info]  2022/11/22 14:26:40.612230 session_raw.go:43: [SQL] SELECT * FROM table WHERE date between ? and ?) ORDER BY id DESC []interface {}{"20220630", "20220702"}`
	sqls := ConvertLogSQLDetails(logid, str)
	if len(sqls) == 0 {
		t.Error("converting error")
	} else {
		sqls.InsertTable(engine)
	}
}
