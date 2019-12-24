package controllers

import (
	"html/template"
	"log"
	"net/http"
)

//ServeViewController serves html file optionally with template
func ServeViewController(file, templateFile string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var view *template.Template
		var parseFilesErr error

		// terminates function if "file" is empty
		if file == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// parse files with template or without template
		if templateFile != "" {
			view, parseFilesErr = template.ParseFiles(templateFile, file)
			if parseFilesErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			view, parseFilesErr = template.ParseFiles(file)
			if parseFilesErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(parseFilesErr)
				return
			}
		}

		// w.Header().Set("Content-Type", "text/html")
		// serve view
		view.Execute(w, nil)
	})
}
