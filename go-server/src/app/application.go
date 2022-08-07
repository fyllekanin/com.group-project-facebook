package app

import (
	"database/sql"
	"github.com/gorilla/mux"
)

type Application struct {
	Db     *sql.DB
	Router *mux.Router
}

func GetNewApplication(router *mux.Router, db *sql.DB) *Application {
	return &Application{
		Db:     db,
		Router: router,
	}
}
