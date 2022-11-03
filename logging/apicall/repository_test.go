package apicall

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestRepository(t *testing.T) {
	driver := "mysql"
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOCAL_DB_USER"), os.Getenv("LOCAL_DB_PASSWORD"), os.Getenv("LOCAL_MYSQL"))

	engine, err := xorm.NewEngine(driver, dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	log := APICall{
		Env:        "test",
		ModuleName: "TEST",
		LogType:    InternalShaingAPI.String(),
		FullURL:    "https://www.daum.net",
		Request:    `{}`,
		Response:   `{}`,
		ErrorMsg:   "error",
		MemberID:   1,
		DonationID: 1,
		OrgID:      1,
		CourierID:  1,
	}

	log.CheckAndMakeTable(engine)

	if err = log.InsertTable(engine); err != nil {
		t.Error(err)
	}
}
