package controllers

import (
	"html/template"
	"net/http"
)

//ServeViewController serves html file optionally with template
func ServeViewController(file, templateFile string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var view *template.Template
		var parseFilesErr error

		// parse files with template or without template
		if templateFile != "" {
			view, parseFilesErr = template.ParseFiles(templateFile, file)
			if parseFilesErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			view, parseFilesErr = template.ParseFiles(file)
			if parseFilesErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		view.Execute(w, nil)
	})
}
