package alerts

import (
	"item-alerts/internal/models"
)

type Repository interface {
	FetchAlerts() ([]models.Alerts, error)
	Match([]string, models.Item) bool
	ProcessAlerts([]models.Alerts, []models.Item)
	ProcessAlert(models.Alerts, []models.Item) error
	RunAlerts()
}
