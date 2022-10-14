package logging

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lib "github.com/beautiful-store/platform-service-library"
	"github.com/labstack/echo"
)

func TestBehavior(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.WriteLog(c)
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}

func TestBehaviorWithParams(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var b *bytes.Buffer

	log := New("testModule").
		WithEnv("dev").
		WithParentService("parentServiceID", "parentServiceName").
		WithMemberID(int64(1)).WithMemberName("member1").WithMemberOrgID(int64(1)).
		WithStack(b)

	log.Begin(c)
	log.WriteLog(c)
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}

func TestBehaviorWithError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.SetError("error message")
	log.WriteLog(c)
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}
func TestBehaviorWithPanic(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.SetPanic("panic message")
	log.WriteLog(c)
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}
func TestBehaviorOutToSNS(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.WriteLog(c)

	topic := os.Getenv("AWSSNS_BEHAVIOR_LOG_TOPIC")
	if topic == "" {
		topic = "arn:aws:sns:region:123456789:test_topic"
	}

	err := log.OutToSNS(topic)
	if err != nil {
		t.Error(err)
	}
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}
