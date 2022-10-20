package logging

import (
	"fmt"
	"io/ioutil"

	"github.com/go-xorm/xorm"
)

func (l *Log) CheckTable(session *xorm.Engine) bool {
	exist, err := session.IsTableExist(l.Context.TableName())
	if err != nil {
		panic(err)
	}

	return exist
}

func (l *Log) CheckAndMakeTable(engine *xorm.Engine) error {
	exist, err := engine.IsTableExist(l.Context.TableName())
	if err != nil {
		return err
	}

	if !exist {
		c, ioErr := ioutil.ReadFile("./ddl.sql")
		if ioErr != nil {
			return ioErr
		}

		query := string(c)
		if _, err = engine.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func (l *Log) InsertTable(engine *xorm.Engine) error {
	if affected, err := engine.Insert(l.Context); err != nil {
		return err
	} else if affected != 1 {
		return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
	}

	return nil
}
