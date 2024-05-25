package handler

import (
	"backend/pkg/itemdata"
	"net/http"
)

type IItemDataHandler struct{}

type ItemDataHandler struct {
	ItemDataService itemdata.IItemData
}

func NewItemDataHandler() *ItemDataHandler {
	return &ItemDataHandler{
		itemdata.NewItemData(),
	}
}

func (handler *ItemDataHandler) ItemDatasController(req *http.Request) {
	handler.ItemDataService.RegisterItem()

}
