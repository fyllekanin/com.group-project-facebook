package product_repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestShouldReturnProduct(t *testing.T) {
	// Given
	var id = 5
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "createdAt", "updatedAt"}).
		AddRow(id, "name", "description", 0, 0, 0)
	mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM products WHERE id = %d LIMIT 1", id)).WillReturnRows(rows)
	var repository = NewProductRepository(db)

	// When
	response, _ := repository.GetProduct(id)

	// Then
	if response.Id != 5 {
		t.Fatalf("Response did not have id 5")
	}
}

func TestShouldReturnProducts(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "createdAt", "updatedAt"}).
		AddRow(1, "name", "description", 0, 0, 0).
		AddRow(2, "name", "description", 0, 0, 0).
		AddRow(3, "name", "description", 0, 0, 0)
	mock.ExpectQuery(fmt.Sprintf("SELECT (.+) FROM products LIMIT %d OFFSET %d", 10, 10)).WillReturnRows(rows)
	var repository = NewProductRepository(db)

	// When
	response, _ := repository.GetProducts(10, 10)

	// Then
	if len(response) != 3 {
		t.Fatalf(fmt.Sprintf("Response has %d rows instead of 3", len(response)))
	}
}
