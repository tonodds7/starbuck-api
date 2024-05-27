package handler

import (
	"backend/models"
	"backend/pkg/itemdata"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type ItemDataHandler struct {
	ItemDataService itemdata.IItemData
}

func NewItemDataHandler() *ItemDataHandler {
	return &ItemDataHandler{
		itemdata.NewItemData(),
	}
}

func (handler *ItemDataHandler) GetAllItem(w http.ResponseWriter, req *http.Request) {
	item, err := handler.ItemDataService.GetAllItem()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
	}

}

func (handler *ItemDataHandler) GetItemDetailById(w http.ResponseWriter, req *http.Request, id string) {

	//convert id type string to id type int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid integer format", http.StatusBadRequest)
		return
	}

	item, err := handler.ItemDataService.GetItemDetailById(idInt)
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
	}

}

func (handler *ItemDataHandler) OrderItem(w http.ResponseWriter, req *http.Request) {

	var bodyBytes []byte
	if req.Body != nil {
		defer req.Body.Close()
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var requestBody models.RequestOrder

	err := json.Unmarshal(bodyBytes, &requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	item, err := handler.ItemDataService.OrderItem(requestBody.ID, requestBody.Quantity)
	if err != nil {
		http.Error(w, "Failed to order items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
	}

}
