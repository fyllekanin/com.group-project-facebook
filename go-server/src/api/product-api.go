package api

import (
	"encoding/json"
	"fmt"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/fyllekanin/go-server/src/entities"
	"github.com/fyllekanin/go-server/src/repositories"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductApi struct {
	application *app.Application
}

func (api *ProductApi) getProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page, _ = strconv.Atoi(params["page"])
	var repository = repositories.NewProductRepository(api.application.Db)

	var productsCount = repository.GetProductsCount()

	var start = (10 * page) - 10
	var products = repositories.NewProductRepository(api.application.Db).GetProducts(start, 10)

	var response = entities.PaginationEntity[entities.ProductEntity]{
		Items:    products,
		Page:     page,
		LastPage: (productsCount / 10) + 1,
	}

	json.NewEncoder(w).Encode(response)
}

func (api *ProductApi) getProduct(w http.ResponseWriter, r *http.Request) {
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

func GetNewProductApi(application *app.Application) *ProductApi {
	var api = &ProductApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/products").Subrouter()

	subRouter.HandleFunc("/page/{page}", api.getProducts).Methods("GET")
	subRouter.HandleFunc("/{id}", api.getProduct).Methods("GET")
	return api
}
