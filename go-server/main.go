package main

import (
	"database/sql"
	"fmt"
	"github.com/fyllekanin/go-server/src/api/product-api"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var db = createDatabaseConnection()
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	var application = app.GetNewApplication(apiRouter, db)
	product_api.GetNewProductApi(application)

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createDatabaseConnection() *sql.DB {
	connStr := "postgres://username:password@localhost/go-server?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
