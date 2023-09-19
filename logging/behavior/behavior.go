package behavior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http/httputil"
	"strconv"

	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"

	"github.com/aws/aws-sdk-go-v2/aws"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/beautiful-store/platform-service-library/aws/sns"
)

const (
	HeaderXRequestID    = "X-Request-ID"
	HeaderXActionID     = "X-Action-ID"
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"
	HeaderContentLength = "Content-Length"
	HeaderTraceID       = "TraceID"
	HeaderServiceID     = "ServiceID"
)

type Log struct {
	Context *logContext
}

type logContext struct {
	ID                int64  `xorm:"id pk autoincr" json:"-"`
	Env               string `json:"env" xorm:"env"`
	ModuleName        string `json:"module_name" xorm:"module_name"`
	TimeUnixNano      int64  `json:"time_unix_nano" xorm:"time_unix_nano"`
	Timestamp         string `json:"timestamp" xorm:"timestamp"`
	TraceID           string `json:"trace_id" xorm:"trace_id"`
	ServiceID         string `json:"service_id" xorm:"service_id"`
	ServiceName       string `json:"service_name" xorm:"service_name"`
	ParentServiceID   string `json:"parent_service_id" xorm:"parent_service_id"`
	ParentServiceName string `json:"parent_service_name" xorm:"parent_service_name"`

	RemoteIP  string `json:"remote_ip" xorm:"remote_ip"`
	URI       string `json:"uri" xorm:"uri"`
	Host      string `json:"host" xorm:"host"`
	Method    string `json:"method" xorm:"method"`
	Path      string `json:"path" xorm:"path"`
	Referer   string `json:"referer" xorm:"referer"`
	UserAgent string `json:"user_agent" xorm:"user_agent"`

	BytesIn  int64 `json:"bytes_in" xorm:"bytes_in"`
	BytesOut int64 `json:"bytes_out" xorm:"bytes_out"`

	Header string `json:"header" xorm:"header"`
	Query  string `json:"query" xorm:"query"`
	Body   string `json:"body" xorm:"body"`

	Status int    `json:"status" xorm:"status"`
	Panic  bool   `json:"panic" xorm:"panic"`
	Error  string `json:"error" xorm:"error"`

	StackTrace string `json:"stack_trace" xorm:"stack_trace"`
	SQLTrace   string `json:"sql_trace" xorm:"sql_trace"`

	Latency int64 `json:"latency" xorm:"latency"`

	MemberID    int64  `json:"member_id" xorm:"member_id"`
	MemberOrgID int64  `json:"member_orgid" xorm:"member_orgid"`
	MemberName  string `json:"member_name" xorm:"member_name"`

	StartTime time.Time     `json:"-" xorm:"-"`
	Stack     *bytes.Buffer `json:"-" xorm:"-"`
	SQL       *bytes.Buffer `json:"-" xorm:"-"`
}

func (*logContext) TableName() string {
	return "behavior_logs"
}

func New(moduleName string) *Log {
	return &Log{
		Context: newLogContext(moduleName),
	}
}

func generator() string {
	return random.String(32)
}

func (l *Log) Begin(c echo.Context) {
	l.Context.TimeUnixNano = time.Now().UTC().UnixNano()
	l.Context.Timestamp = lib.GetDefaultLogLocalDateTimeMilli()
	l.Context.StartTime = time.Now()

	req := c.Request()

	serviceID := c.Request().Header.Get(echo.HeaderXRequestID)
	if serviceID == "" {
		serviceID = generator()
	}
	traceID := c.Request().Header.Get(HeaderTraceID)
	if traceID == "" {
		traceID = serviceID
	}
	parentServiceID := c.Request().Header.Get(HeaderServiceID)

	serviceName := "UNKNOWN"
	for _, r := range c.Echo().Routes() {
		if r.Method == req.Method && r.Path == c.Path() {
			serviceName = r.Name
			break
		}
	}

	realIP := req.RemoteAddr
	if ip := req.Header.Get(HeaderXRealIP); ip != "" {
		realIP = ip
	} else {
		realIP, _, _ = net.SplitHostPort(realIP)
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	bytesIn, _ := strconv.ParseInt(req.Header.Get(HeaderContentLength), 10, 64)

	var queryString string
	if len(req.URL.Query()) > 0 {
		params := make(map[string]interface{}, len(req.URL.Query()))
		for k, v := range req.URL.Query() {
			params[k] = v[0]
		}
		q, _ := lib.Map2Byte(params)
		queryString = string(q)
	}

	var body string
	b, _ := io.ReadAll(req.Body)
	c.Request().Body = io.NopCloser(bytes.NewReader(b))

	if b != nil {
		// body = string(b)
		// bReplaced := passwordRegex.ReplaceAll(b, []byte(`"$1": "*"`))
		var bodyParam interface{}
		d := json.NewDecoder(bytes.NewBuffer(b))
		d.UseNumber()
		if err := d.Decode(&bodyParam); err == nil {
			// body = fmt.Sprintf("%v", bodyParam)
			body, _ = lib.Struct2Json(bodyParam)
		} else {
			body = string(b)
		}
	}

	requestDump, _ := httputil.DumpRequestOut(req, true)

	l.Context.TraceID = traceID
	l.Context.ParentServiceID = parentServiceID
	l.Context.ServiceID = serviceID
	l.Context.ServiceName = serviceName
	l.Context.RemoteIP = realIP
	l.Context.URI = req.RequestURI
	l.Context.Host = req.Host
	l.Context.Method = req.Method
	l.Context.Path = path
	l.Context.Referer = req.Referer()
	l.Context.UserAgent = req.UserAgent()
	l.Context.BytesIn = bytesIn
	l.Context.Query = queryString
	l.Context.Body = body
	l.Context.Header = string(requestDump)
}

func (l *Log) WithMemberID(id int64) *Log {
	l.Context.MemberID = id
	return l
}

func (l *Log) WithMemberName(name string) *Log {
	l.Context.MemberName = name
	return l
}

func (l *Log) WithMemberOrgID(orgid int64) *Log {
	l.Context.MemberOrgID = orgid

	return l
}

func (l *Log) WithStack(b *bytes.Buffer) *Log {
	l.Context.Stack = b

	return l
}

func (l *Log) WithSQL(b *bytes.Buffer) *Log {
	l.Context.SQL = b

	return l
}

func (l *Log) WithEnv(env string) *Log {
	l.Context.Env = env

	return l
}

func (l *Log) SetError(message string) {
	l.Context.Panic = false
	l.Context.Error = message
}

func (l *Log) SetPanic(message string) {
	l.Context.Panic = true
	l.Context.Error = message
}

func (l *Log) WriteLog(c echo.Context) {
	res := c.Response()
	if res != nil {
		l.Context.Status = res.Status
		l.Context.BytesOut = res.Size
	}

	// dumpResponse, _ := httputil.DumpResponse(req.Response, true)
	// l.Body = string(dumpResponse)

	if l.Context.Stack != nil {
		var traces []string
		for {
			trace, err := l.Context.Stack.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			traces = append(traces, string(trace))
		}
		l.Context.StackTrace = strings.Join(traces, "")
	}

	if l.Context.SQL != nil {
		var traces []string
		for {
			trace, err := l.Context.SQL.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			traces = append(traces, string(trace))
		}
		l.Context.SQLTrace = strings.Join(traces, "")
	}
	l.Context.Latency = time.Since(l.Context.StartTime).Milliseconds()

	// l.Context.Stack.Reset()
	// l.Context.SQL.Reset()
}

func (l *Log) Write(c *echo.Context) {
	res := (*c).Response()
	if res != nil {
		l.Context.Status = res.Status
		l.Context.BytesOut = res.Size
	}

	dumpResponse, _ := httputil.DumpResponse((*c).Request().Response, true)
	l.Context.Header = string(dumpResponse)

	// behavior trace
	if l.Context.Stack != nil {
		var traces []string
		for {
			trace, err := l.Context.Stack.ReadBytes('\n')
			if err == io.EOF {
				break
			}

			// traces = append(traces, strings.Trim(string(trace), "\n\t"))
			traces = append(traces, string(trace))
		}
		l.Context.StackTrace = strings.Join(traces, "")
	}

	// sql trace
	if l.Context.SQL != nil {
		var traces []string
		for {
			trace, err := l.Context.SQL.ReadBytes('\n')
			if err == io.EOF {
				break
			}

			// traces = append(traces, strings.Trim(string(trace), "\n\t"))
			traces = append(traces, string(trace))
		}
		l.Context.SQLTrace = strings.Join(traces, "\n")
	}

	l.Context.Latency = time.Since(l.Context.StartTime).Milliseconds()

	l.Context.Stack.Reset()
	l.Context.SQL.Reset()
}

func (l *Log) OutToConsole() {
	fmt.Println(l.Context)
}

func (l *Log) OutToSNS(cfg aws.Config, topic string) error {
	_, err := sns.NewSNS(cfg).WithTopic(topic).Send(sns.Behavior.String(), l.Context)
	if err != nil {
		return err
	}

	return nil
}

func newLogContext(moduleName string) *logContext {
	if moduleName == "" {
		panic("모듈명을 입력하여 주세요")
	}

	return &logContext{
		ModuleName: moduleName,
	}
}
