package product_api

import (
	"encoding/json"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/fyllekanin/go-server/src/entities"
	"github.com/fyllekanin/go-server/src/repositories/product-repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductApi struct {
	application *app.Application
}

func (api *ProductApi) GetProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var page, _ = strconv.Atoi(params["page"])
	var repository = product_repository.NewProductRepository(api.application.Db)

	var productsCount = repository.GetProductsCount()

	var start = (10 * page) - 10
	var products = product_repository.NewProductRepository(api.application.Db).GetProducts(start, 10)

	var response = entities.PaginationEntity[entities.ProductEntity]{
		Items:    products,
		Page:     page,
		LastPage: (productsCount / 10) + 1,
	}

	json.NewEncoder(w).Encode(response)
}

func (api *ProductApi) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id, _ = strconv.Atoi(params["id"])
	var repository = product_repository.NewProductRepository(api.application.Db)

	var entity = repository.GetProduct(id)
	json.NewEncoder(w).Encode(entity)
}

func GetNewProductApi(application *app.Application) *ProductApi {
	var api = &ProductApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/products").Subrouter()

	subRouter.HandleFunc("/page/{page}", api.GetProducts).Methods("GET")
	subRouter.HandleFunc("/{id}", api.GetProduct).Methods("GET")
	return api
}
