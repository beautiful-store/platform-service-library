package apicall

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

type APICall struct {
	ID         int64  `xorm:"id pk autoincr" json:"-"`
	Env        string `xorm:"env"  json:"env"`
	ModuleName string `xorm:"module_name"  json:"moduleName"`
	LogType    string `xorm:"log_type"  json:"logType"`
	FullURL    string `xorm:"full_url"  json:"fullURL"`
	Request    string `xorm:"json 'request'"  json:"request"`
	Response   string `xorm:"json 'response'"  json:"response"`
	ErrorMsg   string `xorm:"error_msg"  json:"errorMsg"`
	DonationID int64  `xorm:"donation_id"  json:"donationID"`
	CourierID  int64  `xorm:"courier_id"  json:"courierID"`
	OrgID      int64  `xorm:"org_id"  json:"orgID"`
	MemberID   int64  `xorm:"member_id"  json:"memberID"`
}

func (*APICall) TableName() string {
	return "api_calls"
}

func (a *APICall) CheckTable(engine *xorm.Engine) bool {
	exist, err := engine.IsTableExist(a.TableName())
	if err != nil {
		panic(err)
	}

	return exist
}

func (a *APICall) CheckAndMakeTable(engine *xorm.Engine) {
	if exist := a.CheckTable(engine); !exist {
		sql := a.sqlCreateTableAPICall()
		if _, err := engine.Exec(sql); err != nil {
			panic(err)
		}
	}
}

func (a *APICall) InsertTable(engine *xorm.Engine) error {
	if affected, err := engine.Insert(a); err != nil {
		return err
	} else if affected != 1 {
		return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
	}

	return nil
}

func (a *APICall) sqlCreateTableAPICall() string {
	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
	(
		id            		INT           NOT NULL auto_increment,
		env           		VARCHAR(20)   NOT NULL,
		module_name   		VARCHAR(60)   NOT NULL,
		log_type      		VARCHAR(20)   NOT NULL,
		full_url      		VARCHAR(200)  NOT NULL,
		request       		JSON          NULL,
		response      		JSON          NULL,
		error_msg     		TEXT          NULL,
		donation_id    		INT           NULL,
		courier_id    		INT           NULL,
		org_id        		INT           NULL,
		member_id     		INT           NULL,
		created_at        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`, a.TableName())
}
