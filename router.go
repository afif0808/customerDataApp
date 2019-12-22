package main

import (
	"customerManagementApp/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", controllers.ServeViewController(
		"../src/customerManagementApp/views/home.html",
		"../src/customerManagementApp/views/template.html"))
	return router
}
