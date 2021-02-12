package controllers

import (
	"forum/sessions"
	"net/http"
)

func InternalError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "500 Internal Error")
		return true
	}
	return false
}
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	user, _ := sessions.GetUser(w, r)
}
