package main

import (
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"item-alerts/internal/alerts"
	"item-alerts/internal/aws"
	"item-alerts/internal/db"
	"item-alerts/internal/items"
	"item-alerts/internal/notifications"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	// Connect to the database
	databaseService := db.NewDatabaseService()
	err = databaseService.Init()
	if err != nil {
		panic("Error initializing database: " + err.Error())
	}

	// Close the database connection when the app exits
	defer func() {
		err := databaseService.Close()
		if err != nil {
			panic("Error closing database: " + err.Error())
		}
	}()

	awsService := aws.NewAWSService(aws.NewAWSRepository())

	notificationService := notifications.NewNotificationService(notifications.NewNotificationRepository(awsService))
	itemService := items.NewItemService(items.NewItemRepository(databaseService))
	alertService := alerts.NewAlertService(alerts.NewAlertRepository(databaseService, itemService, notificationService))

	// Add a cron job to run every night at 9:05 PM
	c := cron.New()
	err = c.AddFunc("5 21 * * *", alertService.RunAlerts)
	if err != nil {
		panic("Error adding cron job: " + err.Error())
	}
	c.Start()
}
