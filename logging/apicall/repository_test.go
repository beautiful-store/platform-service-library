package apicall

import (
	"fmt"
	"os"
	"testing"
	"time"

	lib "github.com/beautiful-store/platform-service-library"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestRepository(t *testing.T) {
	driver := os.Getenv("LOG_DB_DRIVER")
	dbConnection := fmt.Sprintf("%s:%s%s", os.Getenv("LOG_DB_USER"), os.Getenv("LOG_DB_PASSWORD"), os.Getenv("LOG_DB_CONNECTION"))

	engine, err := xorm.NewEngine(driver, dbConnection)
	if err != nil {
		panic(fmt.Errorf("database open error: error: %s", err))
	}

	log := APICall{
		Env:        "test",
		ModuleName: "TEST",
		Timestamp:  time.Now().Format(lib.DateLayout19),
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
