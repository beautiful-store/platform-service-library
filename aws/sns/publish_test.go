package sns

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func TestSendSNS(t *testing.T) {
	topic := os.Getenv("AWSSNS_BEHAVIOR_LOG_TOPIC")
	region := os.Getenv("AWS_DEFAULT_REGION")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAcessKeyID := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAcessKeyID, "")),
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	sns := NewSNS(cfg).WithTopic(topic)

	_, err = sns.Send2(Behavior.String(), "this is test1111")
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = sns.Send2(Behavior.String(), "this is test222")
	if err != nil {
		t.Fatal(err.Error())
	}
}
