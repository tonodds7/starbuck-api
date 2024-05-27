package fake

import (
	"backend/models"

	"github.com/stretchr/testify/mock"
)

type MockItemDataRepository struct {
	mock.Mock
}

func (mock *MockItemDataRepository) GetAllItem() ([]models.Item, error) {
	result := mock.Called()
	return result.Get(0).([]models.Item), result.Error(1)
}

func (mock *MockItemDataRepository) GetItemDetailById(id int) (models.ItemDetail, error) {
	result := mock.Called(id)
	return result.Get(0).(models.ItemDetail), result.Error(1)
}

func (mock *MockItemDataRepository) GetQuantityItemById(id int) (int, int, error) {
	result := mock.Called(id)
	return result.Get(0).(int), result.Get(1).(int), result.Error(2)
}

func (mock *MockItemDataRepository) UpdateItemQuantity(id int, orderQuantity int) (string, error) {
	result := mock.Called(id, orderQuantity)
	return result.Get(0).(string), result.Error(1)
}
