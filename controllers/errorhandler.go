package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"forum/sessions"
	"forum/templ"
	"net/http"
)

func InternalError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError, "500 Internal Error")
		return true
	}
	return false
}
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	user, _ := sessions.GetUser(w, r)
	categories, _ := database.GetCategories()
	templ.ExecTemplate(w, "errorhandler.html", models.PageData{message, categories, user, message})
}
