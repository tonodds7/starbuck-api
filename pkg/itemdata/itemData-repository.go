package itemdata

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
)

type IItemDataRepository interface {
	GetAllItem() ([]Item, error)
}

type itemDataRepository struct {
}

func NewItemDataRepository() IItemDataRepository {
	return &itemDataRepository{}
}

func (repo *itemDataRepository) GetAllItem() ([]Item, error) {
	// Connect to MySQL database
	config := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return []Item{}, err
	}
	defer db.Close()

	//query
	query := "SELECT id,name,price,quantity FROM items"

	var item []Item

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//loop row and append for return
	for rows.Next() {
		var itemRow Item
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.ID, &itemRow.Name, &itemRow.Price, &itemRow.Quantity); err != nil {
			return nil, err
		}
		// Append the item to the slice
		item = append(item, itemRow)
	}

	return item, nil
}
