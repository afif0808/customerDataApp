package main

import (
	"customerManagementApp/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/pelanggan", controllers.ServeViewController(
		"../src/customerManagementApp/views/pelanggan.html",
		"../src/customerManagementApp/views/template.html"))

	// assets serve static files such as css,javascript,image dll.
	router.PathPrefix("/assets/").Handler(controllers.ServeFolderController(
		"/assets/", "../src/customerManagementApp/assets/"))
	return router
}
