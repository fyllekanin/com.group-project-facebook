package api

import (
	"encoding/json"
	"fmt"
	"github.com/fyllekanin/go-server/src/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
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
