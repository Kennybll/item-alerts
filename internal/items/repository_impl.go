package items

import (
	"item-alerts/internal/db"
	"item-alerts/internal/models"
	"time"
)

type ItemRepositoryImpl struct {
	db *db.DatabaseService
}

func NewItemRepository(db *db.DatabaseService) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{
		db: db,
	}
}

func (r *ItemRepositoryImpl) FetchItemsThatStartedToday() ([]models.Item, error) {
	var items []models.Item
	err := r.db.GetDb().Select(&items, "SELECT * FROM items WHERE start_time >= $1 AND start_time < $2", time.Now().Truncate(24*time.Hour), time.Now().Truncate(24*time.Hour).Add(24*time.Hour))
	return items, err
}
