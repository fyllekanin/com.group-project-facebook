package product_api

import (
	"encoding/json"
	"github.com/fyllekanin/go-server/src/app"
	"github.com/fyllekanin/go-server/src/common/error-interface"
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
	var page, err = strconv.Atoi(params["page"])
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: "page needs to be a number",
		})
		return
	}
	var repository = product_repository.NewProductRepository(api.application.Db)
	productsCount, err := repository.GetProductsCount()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: err.Error(),
		})
		return
	}

	var start = (10 * page) - 10
	products, err := product_repository.NewProductRepository(api.application.Db).GetProducts(start, 10)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: err.Error(),
		})
		return
	}

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

	entity, err := repository.GetProduct(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(error_interface.RestError{
			Message: err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func GetApi(application *app.Application) *ProductApi {
	var api = &ProductApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/products").Subrouter()

	subRouter.HandleFunc("/page/{page}", api.GetProducts).Methods("GET")
	subRouter.HandleFunc("/{id}", api.GetProduct).Methods("GET")
	return api
}
