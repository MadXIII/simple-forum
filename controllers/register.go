package controllers

import (
	"forum/database"
	"forum/models"
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
			w.WriteHeader(http.StatusCreated)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if usernameExists {
			data.Data = "Username exists"
			w.WriteHeader(http.StatusCreated)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
			return
		}
		if newUser.Email == "" || !regex.MatchString(newUser.Username) {
			data.Data = "Invalid email"
			w.WriteHeader(http.StatusCreated)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
		}
		if emailExists {
			data.Data = "Email exists"
			w.WriteHeader(http.StatusCreated)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
		}
		if !isValidPass(password) {
			data.Data = "Password must have min 8 characters: at least 1 upper case, 1 lower case, 1 number"
			w.WriteHeader(http.StatusCreated)
			InternalError(w, r, templ.ExecTemplate(w, "register.html", data))
		}
		salt := uuid.NewV4().String()

		hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.MinCost)
		if InternalError(w, r, err) {
			return
		}
		newUser.Salt = salt
		newUser.Hash = hash

		err = database.CreateUser(&newUser)
		if InternalError(w, r, err) {
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
	if r.Method == http.MethodGet {
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