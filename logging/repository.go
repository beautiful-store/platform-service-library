package logging

import (
	"fmt"

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
		sql := sqlCreateTableBehaviorLog()
		if _, err = engine.Exec(sql); err != nil {
			return err
		}

		dsql := sqlCreateTableBehaviorLogDetail()
		if _, err = engine.Exec(dsql); err != nil {
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

	str := l.Context.StackTrace
	if str != "" {
		stacks := ConvertLogDetails(l.Context.ID, str)
		if stacks != nil && len(stacks) > 0 {
			stacks.InsertTable(engine)
		}
	}

	return nil
}

func sqlCreateTableBehaviorLog() string {
	return `CREATE TABLE IF NOT EXISTS behavior_logs
	(
		id                  INT           NOT NULL AUTO_INCREMENT,
		env                 VARCHAR(20)   NOT NULL,
		module_name         VARCHAR(60)   NOT NULL,
		time_unix_nano      BIGINT        NOT NULL,
		timestamp           VARCHAR(60)   NOT NULL,
		service_id	        VARCHAR(60)   NOT NULL,
		service_name        VARCHAR(200)  NOT NULL,
		parent_service_id   VARCHAR(60)   NOT NULL,
		parent_service_name VARCHAR(200)  NOT NULL,
		remote_ip           VARCHAR(30)   NOT NULL,
		uri                 VARCHAR(1000) NOT NULL,
		host                VARCHAR(100)  NOT NULL,
		method              VARCHAR(10)   NOT NULL,
		path                VARCHAR(200)  NOT NULL,
		referer             VARCHAR(500)  NOT NULL,
		user_agent          VARCHAR(500)  NOT NULL,
		bytes_in            INT           NOT NULL,
		bytes_out           INT           NOT NULL,
		header              VARCHAR(1000) NOT NULL,
		query               VARCHAR(1000) NOT NULL,
		body                TEXT          NOT NULL,
		status              SMALLINT		  NOT NULL,
		panic		            TINYINT(1)    NOT NULL,
		error		            VARCHAR(200)  NOT NULL,
		stack_trace         TEXT          NOT NULL,
		latency             INT           NOT NULL,
		member_id           INT           NOT NULL,
		member_name         VARCHAR(200)  NOT NULL,
		member_orgid        INT           NOT NULL,
		created_at          DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
}

func sqlCreateTableBehaviorLogDetail() string {
	return `CREATE TABLE IF NOT EXISTS behavior_log_details
	(
		id						INT           NOT NULL AUTO_INCREMENT,
		log_id 				INT           NOT,
		db		        TINYINT(1)    NOT NULL,
		file          VARCHAR(1000) NOT NULL,
		func          VARCHAR(1000) NOT NULL,
		level         VARCHAR(20)   NOT NULL,
		msg           TEXT          NOT NULL,
		time         	VARCHAR(60)   NOT NULL,
		created_at          DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
}
