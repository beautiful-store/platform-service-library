package behavior

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/labstack/echo/v4"

	lib "github.com/beautiful-store/platform-service-library"
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
	region := os.Getenv("AWS_DEFAULT_REGION")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAcessKeyID := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAcessKeyID, "")),
	)
	if err != nil {
		t.Error(err.Error())
	}
	err = log.OutToSNS(cfg, topic)
	if err != nil {
		t.Error(err)
	}
	log.OutToConsole()
	t.Log(lib.Struct2Json(log.Context))
}
