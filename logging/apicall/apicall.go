package apicall

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/beautiful-store/platform-service-library/aws/sns"

	awssns "github.com/beautiful-store/platform-service-library/aws/sns"
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

func (a *APICall) OutToSNS(cfg aws.Config, topic string) error {
	_, err := awssns.NewSNS(cfg).WithTopic(topic).Send(sns.APICall.String(), a)
	if err != nil {
		return err
	}

	return nil
}
