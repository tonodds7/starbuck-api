package itemdata

import (
	"backend/models"
	"backend/pkg/internal/fake"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewItemData(t *testing.T) {
	service := NewItemData()

	v := reflect.Indirect(reflect.ValueOf(service))

	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}
}

func TestGetAllItemItem(t *testing.T) {
	t.Run("AllItemItemSuccess", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultRepo := []models.Item{
			{
				ID:       1,
				Name:     "item1",
				Price:    1,
				Quantity: 100,
			},
		}

		mockItemDataRepository.On("GetAllItem").Return(mockResultRepo, nil)

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.GetAllItem()

		assert.NoError(t, err)

		expectedResult := []models.Item{
			{
				ID:       1,
				Name:     "item1",
				Price:    1,
				Quantity: 100,
			},
		}
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("AllItemItemError", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultRepo := []models.Item{}

		mockItemDataRepository.On("GetAllItem").Return(mockResultRepo, errors.New("Error_GetAllItem_repo"))

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.GetAllItem()

		assert.ErrorContains(t, err, "Error_GetAllItem_repo")

		expectedResult := []models.Item{}
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestGetItemDetailById(t *testing.T) {
	t.Run("AllItemItemSuccess", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultRepo := models.ItemDetail{
			ID:                    1,
			Name:                  "item1",
			Detail:                "Spain",
			Category:              "good",
			Acidity:               3,
			Body:                  4,
			Processing_Method:     "AA",
			Tasting_Notes:         "BB",
			Complementary_Flavors: "CC",
			Region:                "DD",
			Price:                 150,
		}

		mockItemDataRepository.On("GetItemDetailById", mock.Anything).Return(mockResultRepo, nil)

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.GetItemDetailById(1)

		assert.NoError(t, err)

		expectedResult := models.ItemDetail{
			ID:                    1,
			Name:                  "item1",
			Detail:                "Spain",
			Category:              "good",
			Acidity:               3,
			Body:                  4,
			Processing_Method:     "AA",
			Tasting_Notes:         "BB",
			Complementary_Flavors: "CC",
			Region:                "DD",
			Price:                 150,
		}
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("ItemDetailError", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultRepo := models.ItemDetail{}

		mockItemDataRepository.On("GetItemDetailById", mock.Anything).Return(mockResultRepo, errors.New("Error_GetItemDetailById_repo"))

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.GetItemDetailById(1)

		assert.ErrorContains(t, err, "Error_GetItemDetailById_repo")

		expectedResult := models.ItemDetail{}
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestOrderItem(t *testing.T) {
	t.Run("OrderSuccess", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultGetQty := 100
		mockResultPrice := 500
		mockUpdateResult := "Order Success"

		mockItemDataRepository.On("GetQuantityItemById", mock.Anything).Return(mockResultGetQty, mockResultPrice, nil)
		mockItemDataRepository.On("UpdateItemQuantity", mock.Anything, mock.Anything).Return(mockUpdateResult, nil)

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.OrderItem(1, 20)

		assert.NoError(t, err)

		expectedResult := "Order Success price: 10000"
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("OutOfItem", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultGetQty := 0
		mockResultPrice := 500

		mockItemDataRepository.On("GetQuantityItemById", mock.Anything).Return(mockResultGetQty, mockResultPrice, nil)

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.OrderItem(1, 20)

		assert.NoError(t, err)

		expectedResult := "Out Of Item"
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("ProductsNotEnough ", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultGetQty := 20
		mockResultPrice := 500

		mockItemDataRepository.On("GetQuantityItemById", mock.Anything).Return(mockResultGetQty, mockResultPrice, nil)

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.OrderItem(1, 30)

		assert.NoError(t, err)

		expectedResult := fmt.Sprintf("This Item have %d in stock", mockResultGetQty)
		assert.Equal(t, expectedResult, actualResult)
	})
	t.Run("OrderError", func(t *testing.T) {

		mockItemDataRepository := new(fake.MockItemDataRepository)

		mockResultRepo := 0
		mockResultPrice := 500

		mockItemDataRepository.On("GetQuantityItemById", mock.Anything).Return(mockResultRepo, mockResultPrice, errors.New("Error_OrderItem_repo"))

		service := itemDataService{
			ItemDataRepository: mockItemDataRepository,
		}

		actualResult, err := service.OrderItem(1, 30)

		assert.ErrorContains(t, err, "Error_OrderItem_repo")

		expectedResult := ""
		assert.Equal(t, expectedResult, actualResult)
	})
}
