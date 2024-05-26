package handler

import (
	"backend/pkg/itemdata"
	"encoding/json"
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

func (handler *ItemDataHandler) GetAllItem(w http.ResponseWriter, req *http.Request) {
	item, err := handler.ItemDataService.GetAllItemItem()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
	}

}
