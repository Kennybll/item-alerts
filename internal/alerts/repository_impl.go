package alerts

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"item-alerts/internal/db"
	"item-alerts/internal/items"
	"item-alerts/internal/models"
	"item-alerts/internal/notifications"
	"log"
)

type RepositoryImpl struct {
	db                  *db.DatabaseService
	itemsService        *items.Service
	notificationService *notifications.Service
}

func NewAlertRepository(db *db.DatabaseService, itemsService *items.Service, notificationService *notifications.Service) *RepositoryImpl {
	return &RepositoryImpl{
		db:                  db,
		itemsService:        itemsService,
		notificationService: notificationService,
	}
}

func (r *RepositoryImpl) RunAlerts() {
	alerts, err := r.FetchAlerts()
	if err != nil {
		// Log error
		return
	}

	foundItems, err := r.itemsService.FetchItemsThatStartedToday()
	if err != nil {
		// Log error
		return
	}

	r.ProcessAlerts(alerts, foundItems)
}

// FetchAlerts fetches all alerts from the database
// Users can have multiple alerts
// Let's merge all alerts for a user into a single Alerts struct
// This will make it easier to process alerts later
func (r *RepositoryImpl) FetchAlerts() ([]models.Alerts, error) {
	var alerts []models.Alert
	err := r.db.GetDb().Select(&alerts, "SELECT * FROM alerts WHERE deleted_at IS NULL")

	// Create a map that will store all alerts for each user
	alertsMap := make(map[string][]models.Alert)

	// Loop through all alerts and group them by user ID
	for _, alert := range alerts {
		alertsMap[alert.UserId] = append(alertsMap[alert.UserId], alert)
	}

	// Loop through the map and create an Alerts struct for each user
	var finalAlerts []models.Alerts
	for userId, userAlerts := range alertsMap {
		alerts := make([]string, len(userAlerts))
		for i, alert := range userAlerts {
			alerts[i] = alert.Alert
		}
		finalAlerts = append(finalAlerts, models.Alerts{
			Alerts: alerts,
			UserId: userId,
		})
	}

	return finalAlerts, err
}

func (r *RepositoryImpl) Match(alerts []string, item models.Item) bool {
	// Use fuzzy search to check if the item name contains the alert string
	// This is a more flexible way to match items
	// It is case-insensitive and allows for partial matches
	return fuzzy.FindFold(item.Name, alerts) != nil || fuzzy.FindFold(item.Description, alerts) != nil
}

func (r *RepositoryImpl) ProcessAlerts(alerts []models.Alerts, items []models.Item) {
	// Loop through each alert
	// Use goroutines to process each alert concurrently
	for _, alert := range alerts {
		go func() {
			err := r.ProcessAlert(alert, items)
			if err != nil {
				log.Printf("Error processing alert: " + err.Error())
			}
		}()
	}
}

func (r *RepositoryImpl) ProcessAlert(alert models.Alerts, items []models.Item) error {
	// Find all items that match the alert
	var foundItems []models.Item
	for _, item := range items {
		if r.Match(alert.Alerts, item) {
			foundItems = append(foundItems, item)
		}
	}

	if len(foundItems) > 0 {
		// Send a notification to the user
		return r.notificationService.SendAlert(alert, foundItems)
	}

	return nil
}
