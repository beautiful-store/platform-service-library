package behavior

type logDetail struct {
	ID    int64  `xorm:"id pk autoincr" json:"-"`
	LogID int64  `json:"logID" xorm:"log_id"`
	Level string `json:"level" xorm:"level"`
	SQL   bool   `json:"sql" xorm:"db"`
	File  string `json:"file" xorm:"file"`
	Func  string `json:"func" xorm:"func"`
	Msg   string `json:"msg" xorm:"msg"`
	Time  string `json:"time" xorm:"time"`
}

func (*logDetail) TableName() string {
	return "behavior_log_details"
}

// revive:disable:unexported-return
func newLogDetail(logID int64) *logDetail {
	return &logDetail{
		LogID: logID,
	}
}
