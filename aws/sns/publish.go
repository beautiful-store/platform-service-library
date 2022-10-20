package sns

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// nolint
type snsPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type awssns struct {
	client *sns.Client
	topic  string
}

// revive:disable:unexported-return
func NewSNS(cfg aws.Config) *awssns {
	return &awssns{client: sns.NewFromConfig(cfg)}
}

func (s *awssns) WithTopic(topic string) *awssns {
	if len(topic) == 0 {
		fmt.Println("can't find the topic")
		return nil
	}
	s.topic = topic

	return s
}

func (s *awssns) Send(message string) (*string, error) {
	var messageID *string

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[panic error] can't recover :", err)
		}
	}()

	if s == nil || s.client == nil {
		return messageID, errors.New("can't fild aws sns or client")
	}
	if len(s.topic) == 0 {
		return messageID, errors.New("can't find the topic")
	}
	if len(message) == 0 {
		return messageID, errors.New("There is no message")
	}

	uq := fmt.Sprintf("m%d", time.Now().UnixNano())

	msgPtr := flag.String(uq, message, "")
	topicPtr := flag.String("t", s.topic, "")

	flag.Parse()

	input := &sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	}

	result, err := s.client.Publish(context.TODO(), input)

	if err != nil {
		return messageID, fmt.Errorf("got an error while publishing the message:%v", err)
	}

	return result.MessageId, nil
}
