package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type RepositoryImpl struct {
	cfg config.Config
	ses *ses.Client
	sns *sns.Client
}

func NewAWSRepository() *RepositoryImpl {
	impl := &RepositoryImpl{}
	err := impl.Init()
	if err != nil {
		panic("Error initializing AWS repository: " + err.Error())
	}
	return impl
}

func (r *RepositoryImpl) Init() error {
	// Initialize AWS services
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	r.cfg = cfg
	r.ses = ses.NewFromConfig(cfg)
	r.sns = sns.NewFromConfig(cfg)
	return nil
}

func (r *RepositoryImpl) SendEmail(input *ses.SendTemplatedEmailInput) error {
	_, err := r.ses.SendTemplatedEmail(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) SendSMS(input *sns.PublishInput) error {
	_, err := r.sns.Publish(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
