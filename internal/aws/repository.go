package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Repository interface {
	Init() error
	SendEmail(*ses.SendTemplatedEmailInput) error
	SendSMS(*sns.PublishInput) error
}
