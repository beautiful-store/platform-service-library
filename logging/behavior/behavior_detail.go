package behavior

type logDetail struct {
	ID        int64  `xorm:"id pk autoincr" json:"-"`
	LogID     int64  `json:"logID" xorm:"log_id"`
	ServiceID string `json:"service_id" xorm:"service_id"`
	Level     string `json:"level" xorm:"level"`
	File      string `json:"file" xorm:"file"`
	Func      string `json:"func" xorm:"func"`
	Msg       string `json:"msg" xorm:"msg"`
	Timestamp string `json:"time" xorm:"timestamp"`
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

type logSQLDetail struct {
	ID        int64  `xorm:"id pk autoincr" json:"-"`
	LogID     int64  `json:"logID" xorm:"log_id"`
	ServiceID string `json:"service_id" xorm:"service_id"`
	Level     string `json:"level" xorm:"level"`
	File      string `json:"file" xorm:"file"`
	Func      string `json:"func" xorm:"func"`
	Msg       string `json:"msg" xorm:"msg"`
	Timestamp string `json:"time" xorm:"timestamp"`
}

func (*logSQLDetail) TableName() string {
	return "behavior_log_sql_details"
}

// revive:disable:unexported-return
func newLogSQLDetail(logID int64) *logSQLDetail {
	return &logSQLDetail{
		LogID: logID,
	}
}
