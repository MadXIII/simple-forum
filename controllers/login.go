package controllers

import (
	"forum/database"
	"forum/models"
	"forum/sessions"
	"forum/templ"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, data models.PageData) {
	data.PageTitle = "Login"

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		usernameExists := database.IsUsernameExists(username)
		emailExists := database.IsEmailExists(username)

		if !strings.ContainsRune(username, '@') {
			if !usernameExists {
				data.Data = "Invalid Username"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, templ.ExecTemplate(w, "login.html", data))
				return
			}
			user, err := database.GetUserByUsername(username)
			if InternalError(w, r, err) {
				return
			}
			err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Wrong password"
				w.WriteHeader(http.StatusUnauthorized)
				InternalError(w, r, templ.ExecTemplate(w, "login.html", data))
				return
			}
			err = sessions.CreateSession(user.UserID, w)
			if InternalError(w, r, err) {
				return
			}
		} else {
			if !emailExists {
				data.Data = "Invalid email"
				w.WriteHeader(http.StatusUnprocessableEntity)
				InternalError(w, r, templ.ExecTemplate(w, "login.html", data))
				return
			}
			user, err := database.GetUserByEmail(username)
			if InternalError(w, r, err) {
				return
			}
			err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password+user.Salt))
			if err != nil {
				data.Data = "Wrong password"
				w.WriteHeader(http.StatusUnauthorized)
				InternalError(w, r, templ.ExecTemplate(w, "login.html", data))
				return
			}
			err = sessions.CreateSession(user.UserID, w)
			if InternalError(w, r, err) {
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, templ.ExecTemplate(w, "login.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
