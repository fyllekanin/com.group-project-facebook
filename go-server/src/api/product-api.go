package api

import (
	"encoding/json"
	"fmt"
	"github.com/fyllekanin/go-server/src/entities"
	"github.com/fyllekanin/go-server/src/repositories"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page, _ = strconv.Atoi(params["page"])
	var repository = repositories.NewProductRepository()

	var productsCount = repository.GetProductsCount()

	var start = (10 * page) - 10
	var products = repositories.NewProductRepository().GetProducts(start, 10)

	var response = entities.PaginationEntity[entities.ProductEntity]{
		Items:    products,
		Page:     page,
		LastPage: (productsCount / 10) + 1,
	}

	json.NewEncoder(w).Encode(response)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])

	var response = entities.ProductEntity{
		Id:          id,
		Name:        fmt.Sprintf("Name #%d", id),
		Description: "Cool description",
		Price:       id,
	}

	json.NewEncoder(w).Encode(response)
}

func ProductApi(router *mux.Router) {

	router.HandleFunc("/page/{page}", getProducts).Methods("GET")
	router.HandleFunc("/{id}", getProduct).Methods("GET")

}
