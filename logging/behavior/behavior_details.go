package behavior

import (
	"fmt"
	"strings"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/go-xorm/xorm"
)

type logDetails []*logDetail

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
