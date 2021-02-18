package controllers

import (
	"fmt"
	"forum/database"
	"forum/models"
	"forum/sessions"
	"forum/templ"
	"net/http"
	"regexp"
	"unicode"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//Register - ...
func Register(w http.ResponseWriter, r *http.Request, data models.PageData) {
	data.PageTitle = "Register"
	if r.Method == http.MethodPost {
		regex := regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_.]+[^\^!#\$%&'\@()*+\/=\?\^\n_{\|}~-]@[a-z]{2,}\.[a-zA-Z]{2,6}$`)
		var newUser models.User
		var err error
		newUser.Username = r.FormValue("username")
		newUser.Email = r.FormValue("email")
		password := r.FormValue("password")

		usernameExists := database.IsUsernameExists(newUser.Username)
		emailExists := database.IsEmailExists(newUser.Email)

		if newUser.Username == "" {
			data.Data = "Invalid username"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if usernameExists {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if newUser.Email == "" || !regex.MatchString(newUser.Email) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if emailExists {
			data.Data = "Email exists"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if !isValidPass(password) {
			data.Data = "Password must have min 8 characters: at least 1 upper case, 1 lower case, 1 number"
			w.WriteHeader(http.StatusUnprocessableEntity)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		salt := uuid.NewV4()

		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt.String()), bcrypt.MinCost)
		if InternalError(w, r, err) {
			return
		}
		newUser.Salt = salt.String()
		newUser.Hash = hash

		fmt.Println(newUser.Hash)
		fmt.Println(newUser.Salt)

		err = database.CreateUser(&newUser)
		if InternalError(w, r, err) {
			return
		}
		err = sessions.CreateSession(newUser.UserID, w)
		if InternalError(w, r, err) {
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)

	} else if r.Method == http.MethodGet {
		InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
	} else {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, "405 Method Not Allowed")
	}
}
func isValidPass(pass string) bool {
	var up, low, num bool
	if len(pass) < 8 {
		return false
	}
	for _, r := range pass {
		if unicode.IsUpper(r) {
			up = true
		}
		if unicode.IsLower(r) {
			low = true
		}
		if unicode.IsNumber(r) {
			num = true
		}
	}
	return up && low && num
}
