package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Service struct {
	repo Repository
}

func NewAWSService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Init() {
	s.repo.Init()
}

func (s *Service) SendEmail(input *ses.SendTemplatedEmailInput) error {
	return s.repo.SendEmail(input)
}

func (s *Service) SendSMS(input *sns.PublishInput) error {
	return s.repo.SendSMS(input)
}
