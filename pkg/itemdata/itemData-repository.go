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
	config := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		fmt.Println("sql", err)
		return []models.Item{}, err
	}
	defer db.Close()

	query := "SELECT id,name,price,quantity FROM items"

	var item []models.Item

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("query", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemRow models.Item
		if err := rows.Scan(&itemRow.ID, &itemRow.Name, &itemRow.Price, &itemRow.Quantity); err != nil {
			fmt.Println("Scaner", err)
			return nil, err
		}
		item = append(item, itemRow)
	}

	return item, nil
}

func (repo *itemDataRepository) GetItemDetailById(id int) (models.ItemDetail, error) {
	config := config.LoadConfig()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return models.ItemDetail{}, err
	}
	defer db.Close()
	query := `
        SELECT 
            itm.id, itm.name, itm.price, 
            id.description, id.acidity, id.body, id.processing_method, id.tasting_notes, id.complementary_flavors, id.region,
            ctg.name AS category
        FROM 
            items itm
        JOIN 
            items_detail id ON itm.detail_id = id.detail_id
        JOIN 
            categories ctg ON itm.category_id = ctg.category_id
        WHERE 
            itm.id = ?
    `

	row := db.QueryRow(query, id)

	var itemRow models.ItemDetail
	err = row.Scan(
		&itemRow.ID, &itemRow.Name, &itemRow.Price,
		&itemRow.Detail, &itemRow.Acidity, &itemRow.Body, &itemRow.Processing_Method, &itemRow.Tasting_Notes, &itemRow.Complementary_Flavors, &itemRow.Region,
		&itemRow.Category,
	)
	if err != nil {
		return models.ItemDetail{}, err
	}

	return itemRow, nil
}

func (repo *itemDataRepository) GetQuantityItemById(id int) (int, int, error) {
	config := config.LoadConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT quantity,price FROM items WHERE id = %d", id)

	var item models.Item

	rows, err := db.Query(query)
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var itemRow models.Item

		if err := rows.Scan(&itemRow.Quantity, &itemRow.Price); err != nil {
			return 0, 0, err
		}
		item = itemRow
	}

	return item.Quantity, int(item.Price), nil
}

func (repo *itemDataRepository) UpdateItemQuantity(id int, orderQuantity int) (string, error) {
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
