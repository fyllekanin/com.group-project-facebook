package product_repository

import (
	"database/sql"
	"errors"
	"github.com/fyllekanin/go-server/src/entities"
	_ "github.com/lib/pq"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func (repository *ProductRepository) GetProductsCount() (int, error) {
	statement, err := repository.db.Prepare("SELECT COUNT(*) from products")
	if err != nil {
		log.Println(err)
		return 0, errors.New("failed to prepare statement")
	}
	row, err := statement.Query()
	if err != nil {
		log.Println(err)
		return 0, errors.New("failed to query statement")
	}
	defer row.Close()

	var count int
	for row.Next() {
		row.Scan(&count)
	}

	return count, err
}

func (repository *ProductRepository) GetProducts(start int, limit int) ([]entities.ProductEntity, error) {
	statement, err := repository.db.Prepare("SELECT * FROM products LIMIT ? OFFSET ?")
	if err != nil {
		log.Println(err)
		return []entities.ProductEntity{}, errors.New("failed to prepare statement")
	}
	rows, err := statement.Query(limit, start)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return []entities.ProductEntity{}, errors.New("failed to query statement")
	}

	return scanRows(rows), err
}

func (repository *ProductRepository) GetProduct(id int) (entities.ProductEntity, error) {
	statement, err := repository.db.Prepare("SELECT * FROM products WHERE id = ? LIMIT 1")
	if err != nil {
		log.Println(err)
		return entities.ProductEntity{}, errors.New("failed to prepare statement")
	}
	rows, err := statement.Query(id)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return entities.ProductEntity{}, errors.New("failed to query statement")
	}

	var response = scanRows(rows)
	return response[0], err
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
