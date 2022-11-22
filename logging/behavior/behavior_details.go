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
		for i, stack := range stacks {
			d := newLogDetail(logID)
			if err := lib.String2Struct(stack, &d); err != nil {
				fmt.Println(i, "**err:", err)
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
			sqlTime := strings.TrimSpace(stack[midx+1 : lidx])[:26]
			sql := strings.TrimSpace(stack[lidx+6:])

			d := newLogSQLDetail(logID)
			d.Level = sqlLevel
			d.Time = sqlTime
			d.Msg = sql

			details = append(details, d)
		}
	}

	return details
}

func (stacks logDetails) InsertTable(engine *xorm.Engine) {
	for i, stack := range stacks {
		if affected, err := engine.Insert(stack); err != nil {
			fmt.Println(i, "**stack:", *stack)
			fmt.Println(i, "**err:", err)
			// return err
			continue
		} else if affected != 1 {
			fmt.Println(i, "**err affected rows:", affected)
			// return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
			continue
		}
	}
}

func (sqls logSQLDetails) InsertTable(engine *xorm.Engine) {
	for i, stack := range sqls {
		if affected, err := engine.Insert(stack); err != nil {
			fmt.Println(i, "**stack:", *stack)
			fmt.Println(i, "**err:", err)
			// return err
			continue
		} else if affected != 1 {
			fmt.Println(i, "**err affected rows:", affected)
			// return fmt.Errorf(fmt.Sprintf("affected rows can't be %d", affected))
			continue
		}
	}
}
