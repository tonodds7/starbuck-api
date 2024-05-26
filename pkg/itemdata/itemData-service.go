package itemdata

type IItemData interface {
	GetAllItemItem() ([]Item, error)
}

type itemDataService struct {
	itemDataRepository IItemDataRepository
}

func NewItemData() IItemData {
	return &itemDataService{
		itemDataRepository: NewItemDataRepository(),
	}
}

func (service *itemDataService) GetAllItemItem() ([]Item, error) {
	data, err := service.itemDataRepository.GetAllItem()
	if err != nil {
		return []Item{}, err
	}

	return data, nil

}
