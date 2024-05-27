package itemdata

import (
	"backend/models"
	"fmt"
)

type IItemData interface {
	GetAllItem() ([]models.Item, error)
	GetItemDetailById(id int) (models.ItemDetail, error)
	OrderItem(id int, quantity int) (string, error)
}

type itemDataService struct {
	ItemDataRepository IItemDataRepository
}

func NewItemData() IItemData {
	return &itemDataService{
		ItemDataRepository: NewItemDataRepository(),
	}
}

func (service *itemDataService) GetAllItem() ([]models.Item, error) {
	data, err := service.ItemDataRepository.GetAllItem()
	if err != nil {
		return []models.Item{}, err
	}

	return data, nil

}

func (service *itemDataService) GetItemDetailById(id int) (models.ItemDetail, error) {
	data, err := service.ItemDataRepository.GetItemDetailById(id)
	if err != nil {
		return models.ItemDetail{}, err
	}

	return data, nil

}

func (service *itemDataService) OrderItem(id int, quantity int) (string, error) {

	quantityInStock, price, err := service.ItemDataRepository.GetQuantityItemById(id)
	if err != nil {
		return "", err
	}

	if quantityInStock == 0 {
		return "Out Of Item", nil
	}

	if quantityInStock >= quantity {
		quantityForUpdate := quantityInStock - quantity
		result, err := service.ItemDataRepository.UpdateItemQuantity(id, quantityForUpdate)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v price: %d", result, quantity*price), nil

	} else {
		return fmt.Sprintf("This Item have %d in stock", quantityInStock), nil
	}
}
