package fake

import "github.com/stretchr/testify/mock"

type MockItemDataRepository struct {
	mock.Mock
}

func (mock *MockItemDataRepository) GetAllItemItem() ([]models.Item, error) {
	result := mock.Called()

	return result.Get(0).([]Item), result.Error(1)

}
