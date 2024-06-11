package alerts

import (
	"item-alerts/internal/models"
)

type Service struct {
	repo Repository
}

func NewAlertService(alertRepository Repository) *Service {
	return &Service{
		repo: alertRepository,
	}
}

func (s *Service) FetchAlerts() ([]models.Alerts, error) {
	return s.repo.FetchAlerts()
}

func (s *Service) Match(alerts []string, item models.Item) bool {
	return s.repo.Match(alerts, item)
}

func (s *Service) ProcessAlerts(alerts []models.Alerts, items []models.Item) {
	s.repo.ProcessAlerts(alerts, items)
}

func (s *Service) ProcessAlert(alerts models.Alerts, items []models.Item) error {
	return s.repo.ProcessAlert(alerts, items)
}

func (s *Service) RunAlerts() {
	s.repo.RunAlerts()
}
