package controllers

import (
	"forum/sessions"
	"net/http"
)

//Logout - ...
func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
