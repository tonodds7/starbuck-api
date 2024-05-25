package itemdata

import "fmt"

type IItemData interface {
	RegisterItem() error
}

type itemDataService struct {
}

func NewItemData() IItemData {
	return &itemDataService{}
}

func (service *itemDataService) RegisterItem() error {
	fmt.Println("TONNNNNN")
	return nil

}
