package repositories

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

func (repository *ProductRepository) GetProducts(start int, limit int) []entities.ProductEntity {
	rows, err := repository.db.Query(fmt.Sprintf("SELECT * FROM products LIMIT %d OFFSET %d", limit, start))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var response []entities.ProductEntity
	for rows.Next() {
		var r entities.ProductEntity
		err := rows.Scan(&r.Id, &r.Name, &r.Description, &r.Price, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, r)
	}
	return response
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

func NewProductRepository() *ProductRepository {
	connStr := "postgres://username:password@localhost/go-server?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &ProductRepository{
		db: db,
	}
}
