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

	router.Handle("/printAlamat", controllers.ServeViewController(
		"../src/customerManagementApp/views/printAlamat.html", ""))

	// assets serve static files such as css,javascript,image dll.
	router.PathPrefix("/assets/").Handler(controllers.ServeFolderController(
		"/assets/", "../src/customerManagementApp/assets/"))

	router.PathPrefix("/spaViews/").Handler(controllers.ServeFolderController(
		"/spaViews/", "../src/customerManagementApp/spaViews/"))

	return router
}
