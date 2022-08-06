package main

import (
	"fmt"
	"github.com/fyllekanin/go-server/src/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/products/page/{page}", api.GetProducts).Methods("GET")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
