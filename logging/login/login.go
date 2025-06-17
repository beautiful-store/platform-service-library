package login

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/beautiful-store/platform-service-library/aws/sns"
)

type Login struct {
	ID         int64  `xorm:"id pk autoincr" json:"-"`
	Env        string `xorm:"env"  json:"env"`
	ModuleName string `xorm:"module_name"  json:"moduleName"`
	Timestamp  string `xorm:"timestamp" json:"timestamp"`
	LogType    string `xorm:"log_type"  json:"logType"`
	MemberID   int64  `xorm:"member_id"  json:"memberID"`
}

func (*Login) TableName() string {
	return "login_logs"
}

func (a *Login) OutToSNS(cfg aws.Config, topic string) error {
	_, err := sns.NewSNS(cfg).WithTopic(topic).Send2(sns.LogIn.String(), a)
	if err != nil {
		return err
	}

	return nil
}
