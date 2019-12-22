package controllers

import "net/http"

// ServeFolderController serve static files in a folder
func ServeFolderController(stripPrefix, folderPath string) http.Handler {
	folder := http.Dir(folderPath)
	fileServer := http.FileServer(folder)

	return http.StripPrefix(stripPrefix, fileServer)
}
