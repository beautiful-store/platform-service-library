package behavior

import (
	"fmt"
	"strings"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/go-xorm/xorm"
)

type logDetails []*logDetail
type logSQLDetails []*logSQLDetail

func ConvertLogDetail(s string) logDetails {
	details := make([]*logDetail, 0)

	stacks := strings.Split(s, "\n")
	if len(stacks) > 0 {
		for i, stack := range stacks {
			if stack == "" {
				continue
			}
			d := newLogDetail(int64(i))
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

// revive:disable:unexported-return
func ConvertLogDetails(logID int64, s string) logDetails {
	details := make([]*logDetail, 0)

	stacks := strings.Split(s, "\n")
	if len(stacks) > 0 {
		for _, stack := range stacks {
			if stack == "" {
				continue
			}
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
			d := newLogSQLDetail(logID)

			fidx := strings.Index(stack, "[")
			midx := strings.Index(stack, "]")
			lidx := strings.Index(stack, "[SQL]")
			if fidx < 0 || midx < 0 || lidx < 0 || len(stack) < lidx+6 {
				if fidx < 0 || midx < 0 || lidx < 0 {
					continue
				}
				d.Msg = stack
				fmt.Println("[ERROR]SQL 메세지 오류입니다.", stack)
			} else {
				serviceID := strings.TrimSpace(stack[:fidx])
				sqlLevel := strings.TrimSpace(stack[fidx+1 : midx])
				s := strings.TrimSpace(stack[midx+1 : lidx])
				bidx := strings.LastIndex(s, " ")
				sqlTime := strings.ReplaceAll(s[:bidx], "/", "-")
				sqlFile := strings.TrimSpace(s[bidx+1:])
				sql := strings.TrimSpace(stack[lidx+6:])

				d.ServiceID = serviceID
				d.Level = sqlLevel
				d.File = sqlFile
				d.Timestamp = sqlTime
				d.Msg = sql
			}

			details = append(details, d)
		}
	}

	return details
}

func (stack logDetail) InsertOne(engine *xorm.Engine) {
	if affected, err := engine.Insert(stack); err != nil {
		fmt.Println("[error]", err, "/n", stack)
	} else if affected != 1 {
		fmt.Println("[error]affected rows:", affected)
	}
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
