package logging

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

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
}

func TestBehaviorWithParams(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var b *bytes.Buffer

	log := New("testModule").
		WithParentService("parentServiceID", "parentServiceName").
		WithMemberID(int64(1)).WithMemberName("member1").WithMemberOrgID(int64(1)).
		WithStack(b)

	log.Begin(c)
	log.WriteLog(c)
	log.OutToConsole()
}

func TestBehaviorWithError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.WithError("error message").WriteLog(c)
	log.OutToConsole()
}
func TestBehaviorWithPanic(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log := New("testModule")

	log.Begin(c)
	log.WithPanic("panic message").WriteLog(c)
	log.OutToConsole()
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
}
