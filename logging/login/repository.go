package login

import (
	"fmt"

	"github.com/go-xorm/xorm"

	lib "github.com/beautiful-store/platform-service-library"
)

func (a *Login) CheckTable(engine *xorm.Engine) bool {
	exist, err := engine.IsTableExist(a.TableName())
	if err != nil {
		panic(err)
	}

	return exist
}

func (a *Login) CheckAndMakeTable(engine *xorm.Engine) {
	if exist := a.CheckTable(engine); !exist {
		sql := a.sqlCreateTable()
		if _, err := engine.Exec(sql); err != nil {
			panic(err)
		}
	}
}

func (a *Login) InsertTable(engine *xorm.Engine) error {
	if a.Timestamp == "" {
		a.Timestamp = lib.GetDefaultLogLocalDateTimeMilli()
	}
	if affected, err := engine.Insert(a); err != nil {
		return err
	} else if affected != 1 {
		return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
	}

	return nil
}

func (a *Login) sqlCreateTable() string {
	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
	(
		id            		INT           NOT NULL auto_increment,
		env           		VARCHAR(20)   NOT NULL,
		module_name   		VARCHAR(60)   NOT NULL,
		timestamp     		VARCHAR(60)   NOT NULL,		
		log_type      		VARCHAR(20)   NOT NULL,
		member_id     		INT           NULL,
		created_at        DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		INDEX idx_members_created_at (created_at DESC, member_id asc)
	);`, a.TableName())
}
