package items

import "item-alerts/internal/models"

type Service struct {
	repo Repository
}

func NewItemService(itemRepository Repository) *Service {
	return &Service{
		repo: itemRepository,
	}
}

func (s *Service) FetchItemsThatStartedToday() ([]models.Item, error) {
	return s.repo.FetchItemsThatStartedToday()
}
