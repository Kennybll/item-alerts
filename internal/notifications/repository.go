package notifications

import (
	"item-alerts/internal/models"
)

type Repository interface {
	SendAlert(alert models.Alerts, items []models.Item) error
	SendAlertEmail(email string, items []models.Item) error
	SendAlertSMS(phone string, items []models.Item) error
}
