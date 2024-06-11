package items

import "item-alerts/internal/models"

type Repository interface {
	FetchItemsThatStartedToday() ([]models.Item, error)
}
