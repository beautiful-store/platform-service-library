package apicall

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"

	lib "github.com/beautiful-store/platform-service-library"
)

func (a *APICall) CheckTable(engine *xorm.Engine) bool {
	exist, err := engine.IsTableExist(a.TableName())
	if err != nil {
		panic(err)
	}

	return exist
}

func (a *APICall) CheckAndMakeTable(engine *xorm.Engine) {
	if exist := a.CheckTable(engine); !exist {
		sql := a.sqlCreateTable()
		if _, err := engine.Exec(sql); err != nil {
			panic(err)
		}
	}
}

func (a *APICall) InsertTable(engine *xorm.Engine) error {
	if a.Timestamp == "" {
		a.Timestamp = time.Now().Local().Format(lib.DateLayout19)
	}
	if affected, err := engine.Insert(a); err != nil {
		return err
	} else if affected != 1 {
		return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
	}

	return nil
}

func (a *APICall) sqlCreateTable() string {
	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
	(
		id            		INT           NOT NULL auto_increment,
		env           		VARCHAR(20)   NOT NULL,
		module_name   		VARCHAR(60)   NOT NULL,
		timestamp     		VARCHAR(60)   NOT NULL,		
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
