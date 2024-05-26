package itemdata

import (
	"backend/models"
	"backend/pkg/internal/fake"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
		
		mockItemDataRepository.On("GetAllItemItem"),Return(mockResultRepo,nil)

		service := itemDataService{
			itemDataRepository: mockItemDataRepository,
		}

		actualResult,err :=service.GetAllItemItem()

		assert.NoError(t,err)
	})
}
