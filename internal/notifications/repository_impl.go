package notifications

import (
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"item-alerts/internal/aws"
	"item-alerts/internal/models"
)

type RepositoryImpl struct {
	awsService *aws.Service
}

func NewNotificationRepository(awsService *aws.Service) *RepositoryImpl {
	return &RepositoryImpl{
		awsService: awsService,
	}
}

func (r *RepositoryImpl) SendAlert(alert models.Alerts, items []models.Item) error {
	return nil
}

func (r *RepositoryImpl) SendAlertEmail(email string, items []models.Item) error {
	return r.awsService.SendEmail(&ses.SendTemplatedEmailInput{})
}

func (r *RepositoryImpl) SendAlertSMS(phone string, items []models.Item) error {
	return r.awsService.SendSMS(&sns.PublishInput{})
}
