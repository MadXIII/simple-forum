package router

import (
	"forum/controllers"
	"forum/database"
	"forum/models"
	"forum/sessions"
	"net/http"
)

func PageDataMiddleWare(loginrequired bool, handler func(http.ResponseWriter, *http.Request, models.PageData)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := sessions.GetUser(w, r)
		if controllers.InternalError(w, r, err) {
			return
		}
		categories, err := database.GetCategories()
		if controllers.InternalError(w, r, err) {
			return
		}
		data := models.PageData{
			PageTitle:  "",
			Categories: categories,
			User:       user,
			Data:       nil,
		}
		if loginrequired && user.UserID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handler(w, r, data)
	}
}
