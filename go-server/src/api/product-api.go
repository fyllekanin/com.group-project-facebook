package api

import (
	"encoding/json"
	"fmt"
	"github.com/fyllekanin/go-server/src/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page, _ = strconv.Atoi(params["page"])
	var products []entities.ProductEntity
	var start = (10 * page) - 10

	for i := start; i < (start + 10); i++ {
		products = append(products, entities.ProductEntity{
			Id:          i,
			Name:        fmt.Sprintf("Name #%d", i),
			Description: "Cool description",
			Price:       i,
		})
	}

	var response = entities.PaginationEntity[entities.ProductEntity]{
		Items:    products,
		Page:     page,
		LastPage: 53,
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
