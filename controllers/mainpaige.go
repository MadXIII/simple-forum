package controllers

import (
	"net/http"
)

func MainPaige(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound, "404 Not Found")
		}
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
