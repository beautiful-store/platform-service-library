package logging

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
