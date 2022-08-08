package product_repository

import (
	"database/sql"
	"fmt"
	"github.com/fyllekanin/go-server/src/entities"
	_ "github.com/lib/pq"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func (repository *ProductRepository) GetProductsCount() int {
	var count int
	row, _ := repository.db.Query("SELECT COUNT(*) from products")
	defer row.Close()

	for row.Next() {
		row.Scan(&count)
	}

	return count
}

func (repository *ProductRepository) GetProducts(start int, limit int) []entities.ProductEntity {
	rows, err := repository.db.Query(fmt.Sprintf("SELECT * FROM products LIMIT %d OFFSET %d", limit, start))
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	return scanRows(rows)
}

func (repository *ProductRepository) GetProduct(id int) entities.ProductEntity {
	rows, err := repository.db.Query(fmt.Sprintf("SELECT * FROM products WHERE id = %d LIMIT 1", id))
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var response = scanRows(rows)
	return response[0]
}

func scanRows(rows *sql.Rows) []entities.ProductEntity {
	var response []entities.ProductEntity
	for rows.Next() {
		var item entities.ProductEntity
		rows.Scan(&item.Id, &item.Name, &item.Description, &item.Price, &item.CreatedAt, &item.UpdatedAt)
		response = append(response, item)
	}
	return response
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
