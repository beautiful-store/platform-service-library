package aws

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// type snsPublishAPI interface {
// 	Publish(ctx context.Context,
// 		params *sns.PublishInput,
// 		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
// }

type awssns struct {
	client *sns.Client
	topic  string
}

// revive:disable:unexported-return
func NewSNS() *awssns {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("configuration error, ", err.Error())
		return nil
	}

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

	if s.client == nil {
		return messageID, errors.New("can't find the client")
	}
	if len(s.topic) == 0 {
		return messageID, errors.New("can't find the topic")
	}
	if len(message) == 0 {
		return messageID, errors.New("There is no message")
	}

	msgPtr := flag.String("message", message, "")
	topicPtr := flag.String("topic", s.topic, "")

	flag.Parse()

	input := &sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	}

	result, err := s.client.Publish(context.TODO(), input)
	if err != nil {
		fmt.Println("Got an error publishing the message:")
		fmt.Println(err)
		return messageID, err
	}

	return result.MessageId, nil
}
