package login

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	lib "github.com/beautiful-store/platform-service-library"
)

func TestOutToSNS(t *testing.T) {
	log := &Login{
		Env:        "test",
		ModuleName: "test",
		LogType:    "",
		MemberID:   1,
	}

	topic := os.Getenv("AWSSNS_LOG_TOPIC")
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

	t.Log(lib.Struct2Json(&log))
}
