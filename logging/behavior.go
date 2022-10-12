package logging

import (
	"bytes"
	"io"
	"net"
	"net/http/httputil"
	"os"
	"strconv"

	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	lib "github.com/beautiful-store/platform-service-library"
)

var (
	layout = time.RFC3339
)

const (
	HeaderXRequestID    = "X-Request-ID"
	HeaderXActionID     = "X-Action-ID"
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"
	HeaderContentLength = "Content-Length"
)

type Log struct {
	Stack   *bytes.Buffer
	Context *logContext
}

type logContext struct {
	ModuleName        string `json:"module_name"`
	TimeUnixNano      int64  `json:"time_unix_nano"`
	Timestamp         string `json:"timestamp"`
	ServiceID         string `json:"service_id"`
	ParentServiceID   string `json:"parent_service_id"`
	ParentServiceName string `json:"parent_service_name"`

	RemoteIP  string `json:"remote_ip"`
	Uri       string `json:"uri"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Referer   string `json:"referer"`
	UserAgent string `json:"user_agent"`

	BytesIn  int64 `json:"bytes_in"`
	BytesOut int64 `json:"bytes_out"`

	Header string `json:"header"`
	Query  string `json:"query"`
	Form   string `json:"form"`

	Status int    `json:"status"`
	Panic  bool   `json:"panic"`
	Error  string `json:"error"`

	Body       string `json:"body"`
	StackTrace string `json:"stack_trace"`

	Latency int64 `json:"latency"`

	MemberID    int64  `json:"member_id"`
	MemberOrgID int64  `json:"member_orgid"`
	MemberName  string `json:"member_name"`

	StartTime time.Time `json:"-"`
}

func New(moduleName string, options ...func(*Log)) *Log {
	l := &Log{
		Context: newLogContext(moduleName),
	}

	for _, o := range options {
		if o != nil {
			o(l)
		}
	}

	return l
}

func StackTrace(b *bytes.Buffer) func(*Log) {
	logger := logrus.New()

	mw := io.MultiWriter(os.Stdout, b)
	logger.SetOutput(mw)

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)

	return func(l *Log) {
		l.Stack = b
	}
}

func (l *Log) WithParentService(serviceID string, serviceName string) *Log {
	l.Context.ParentServiceID = serviceID
	l.Context.ParentServiceName = serviceName

	return l
}

func (l *Log) Begin(c echo.Context) {
	l.Context.TimeUnixNano = time.Now().UTC().UnixNano()
	l.Context.Timestamp = time.Now().Format(layout)
	l.Context.StartTime = time.Now()

	req := c.Request()

	realIP := req.RemoteAddr
	if ip := req.Header.Get(HeaderXForwardedFor); ip != "" {
		realIP = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get(HeaderXRealIP); ip != "" {
		realIP = ip
	} else {
		realIP, _, _ = net.SplitHostPort(realIP)
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	bytesIn, _ := strconv.ParseInt(req.Header.Get(HeaderContentLength), 10, 64)

	qparams := ""
	if len(req.URL.Query()) > 0 {
		params := make(map[string]interface{}, len(req.URL.Query()))
		for k, v := range req.URL.Query() {
			params[k] = v[0]
		}
		q, _ := lib.Map2Byte(params)
		qparams = string(q)
	}

	pparams := ""
	if len(c.ParamNames()) > 0 {
		params := make(map[string]interface{}, len(c.ParamNames()))
		for _, name := range c.ParamNames() {
			params[name] = c.Param(name)
		}
		p, _ := lib.Map2Byte(params)
		pparams = string(p)
	}

	requestDump, _ := httputil.DumpRequestOut(req, true)

	l.Context.RemoteIP = realIP
	l.Context.Uri = req.RequestURI
	l.Context.Host = req.Host
	l.Context.Method = req.Method
	l.Context.Path = path
	l.Context.Referer = req.Referer()
	l.Context.UserAgent = req.UserAgent()
	l.Context.BytesIn = bytesIn
	l.Context.Query = qparams
	l.Context.Form = pparams
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

func (l *Log) WriteLog(c echo.Context) {
	res := c.Response()
	if res != nil {
		l.Context.ServiceID = res.Header().Get(echo.HeaderXRequestID)
		l.Context.Status = res.Status
		l.Context.BytesOut = res.Size
	}

	// dumpResponse, _ := httputil.DumpResponse(c.Request().Response, true)
	// l.Body = string(dumpResponse)

	var traces []string
	for {
		trace, err := l.Stack.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		traces = append(traces, string(trace))
	}
	l.Context.StackTrace = strings.Join(traces, ",")
	l.Context.Latency = time.Now().Sub(l.Context.StartTime).Milliseconds()
}

func newLogContext(moduleName string) *logContext {
	if moduleName == "" {
		panic("모듈명을 입력하여 주세요")
	}

	return &logContext{
		ModuleName: moduleName,
	}
}
