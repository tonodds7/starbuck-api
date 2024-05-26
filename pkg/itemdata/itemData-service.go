package itemdata

import "backend/models"

type IItemData interface {
	GetAllItemItem() ([]models.Item, error)
}

type itemDataService struct {
	itemDataRepository IItemDataRepository
}

func NewItemData() IItemData {
	return &itemDataService{
		itemDataRepository: NewItemDataRepository(),
	}
}

func (service *itemDataService) GetAllItemItem() ([]models.Item, error) {
	data, err := service.itemDataRepository.GetAllItem()
	if err != nil {
		return []models.Item{}, err
	}

	return data, nil

}
