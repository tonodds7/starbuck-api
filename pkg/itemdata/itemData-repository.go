package itemdata

import (
	"backend/internal/config"
	"backend/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type IItemDataRepository interface {
	GetAllItem() ([]models.Item, error)
	GetItemDetailById(id int) (models.ItemDetail, error)
	GetQuantityItemById(id int) (int, int, error)
	UpdateItemQuantity(id int, orderQuantity int) (string, error)
}

type itemDataRepository struct {
}

func NewItemDataRepository() IItemDataRepository {
	return &itemDataRepository{}
}

func (repo *itemDataRepository) GetAllItem() ([]models.Item, error) {
	// Connect to MySQL database
	config := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		fmt.Println("sql", err)
		return []models.Item{}, err
	}
	defer db.Close()

	//query
	query := "SELECT id,name,price,quantity FROM items"

	var item []models.Item

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("query", err)
		return nil, err
	}
	defer rows.Close()

	//loop row and append for return
	for rows.Next() {
		var itemRow models.Item
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.ID, &itemRow.Name, &itemRow.Price, &itemRow.Quantity); err != nil {
			fmt.Println("Scaner", err)
			return nil, err
		}
		// Append the item to the slice
		item = append(item, itemRow)
	}

	return item, nil
}

func (repo *itemDataRepository) GetItemDetailById(id int) (models.ItemDetail, error) {
	// Connect to MySQL database
	config := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return models.ItemDetail{}, err
	}
	defer db.Close()

	//query itemId
	query := fmt.Sprintf("SELECT id,name,price,detail_id,category_id FROM items WHERE id = %d ", id)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("0", err)

		return models.ItemDetail{}, err
	}
	defer rows.Close()

	//loop row and append for return
	var itemRow models.ItemDetail

	var ItemDetailId int
	var CategoryId int
	for rows.Next() {
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.ID, &itemRow.Name, &itemRow.Price, &ItemDetailId, &CategoryId); err != nil {
			return models.ItemDetail{}, err
		}
	}

	//query itemdetail
	query = fmt.Sprintf("SELECT description,acidity,body,processing_method,tasting_notes,complementary_flavors,region FROM items_detail WHERE detail_id = %d ", ItemDetailId)

	rows, err = db.Query(query)
	if err != nil {
		return models.ItemDetail{}, err
	}
	defer rows.Close()

	for rows.Next() {
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.Detail, &itemRow.Acidity, &itemRow.Body, &itemRow.Processing_Method, &itemRow.Tasting_Notes, &itemRow.Complementary_Flavors, &itemRow.Region); err != nil {
			return models.ItemDetail{}, err
		}
	}

	//query itemcategories
	query = fmt.Sprintf("SELECT name FROM categories WHERE category_id = %d ", CategoryId)
	rows, err = db.Query(query)
	if err != nil {
		return models.ItemDetail{}, err
	}
	defer rows.Close()

	for rows.Next() {
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.Category); err != nil {
			return models.ItemDetail{}, err
		}
	}
	return itemRow, nil

}

func (repo *itemDataRepository) GetQuantityItemById(id int) (int, int, error) {
	// Connect to MySQL database
	config := config.LoadConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()

	//query
	query := fmt.Sprintf("SELECT quantity,price FROM items WHERE id = %d", id)

	var item models.Item

	rows, err := db.Query(query)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	//loop row and append for return
	for rows.Next() {
		var itemRow models.Item
		// Scan the columns into the struct fields
		if err := rows.Scan(&itemRow.Quantity, &itemRow.Price); err != nil {
			return 0, 0, err
		}
		// Append the item to the slice
		item = itemRow
	}

	return item.Quantity, int(item.Price), nil
}

func (repo *itemDataRepository) UpdateItemQuantity(id int, orderQuantity int) (string, error) {

	// Connect to MySQL database
	config := config.LoadConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return "", err
	}
	defer db.Close()

	updateQuery := fmt.Sprintf("UPDATE items SET quantity = %d WHERE id = %d", orderQuantity, id)
	_, err = db.Exec(updateQuery)
	if err != nil {
		return "", err
	}

	return "Order Succes", nil

}
