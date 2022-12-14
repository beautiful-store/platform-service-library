package behavior

import (
	"fmt"
	"strings"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/go-xorm/xorm"
)

type logDetails []*logDetail
type logSQLDetails []*logSQLDetail

// revive:disable:unexported-return
func ConvertLogDetails(logID int64, s string) logDetails {
	details := make([]*logDetail, 0)

	stacks := strings.Split(s, "\n")
	if len(stacks) > 0 {
		for _, stack := range stacks {
			d := newLogDetail(logID)
			if err := lib.String2Struct(stack, &d); err != nil {
				fmt.Println("[error]", err)
				continue
			}

			fmt.Println("ok:", d)
			details = append(details, d)
		}
	}

	return details
}

func ConvertLogSQLDetails(logID int64, s string) logSQLDetails {
	details := make([]*logSQLDetail, 0)

	stacks := strings.Split(s, "\n")
	if len(stacks) > 0 {
		for _, stack := range stacks {
			fidx := strings.Index(stack, "[")
			midx := strings.Index(stack, "]")
			lidx := strings.Index(stack, "[SQL]")
			if fidx < 0 || midx < 0 || lidx < 0 {
				fmt.Println("[ERROR]SQL 메세지 오류입니다.", stack)
				continue
			}
			sqlLevel := strings.TrimSpace(stack[fidx+1 : midx])
			s := strings.TrimSpace(stack[midx+1 : lidx])
			bidx := strings.LastIndex(s, " ")
			sqlTime := strings.ReplaceAll(s[:bidx], "/", "-")
			sqlFile := strings.TrimSpace(s[bidx+1:])
			sql := strings.TrimSpace(stack[lidx+6:])

			d := newLogSQLDetail(logID)
			d.Level = sqlLevel
			d.File = sqlFile
			d.Timestamp = sqlTime
			d.Msg = sql

			details = append(details, d)
		}
	}

	return details
}

func (stacks logDetails) InsertTable(engine *xorm.Engine) {
	for i, stack := range stacks {
		if affected, err := engine.Insert(stack); err != nil {
			fmt.Println("[error]", err, "/n", *stack)
			continue
		} else if affected != 1 {
			fmt.Println(i, "[error]affected rows:", affected)
			continue
		}
	}
}

func (sqls logSQLDetails) InsertTable(engine *xorm.Engine) {
	for i, stack := range sqls {
		if affected, err := engine.Insert(stack); err != nil {
			fmt.Println("[error]", err, "/n", *stack)
			continue
		} else if affected != 1 {
			fmt.Println(i, "[error]affected rows:", affected)
			continue
		}
	}
}
