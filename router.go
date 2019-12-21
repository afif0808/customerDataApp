package main

import (
	"customerManagementApp/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", controllers.ServeViewController("/", ""))
	return router
}
