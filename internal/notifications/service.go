package notifications

import (
	"item-alerts/internal/models"
)

type Service struct {
	repo Repository
}

func NewNotificationService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SendAlert(alert models.Alerts, items []models.Item) error {
	return s.repo.SendAlert(alert, items)
}

func (s *Service) SendAlertEmail(email string, items []models.Item) error {
	return s.repo.SendAlertEmail(email, items)
}

func (s *Service) SendAlertSMS(phone string, items []models.Item) error {
	return s.repo.SendAlertSMS(phone, items)
}
