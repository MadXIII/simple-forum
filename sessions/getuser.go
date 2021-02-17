package sessions

import (
	"forum/database"
	"forum/models"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User

	ck, err := r.Cookie("session")
	if err != nil {
		return user, err
	}
	session, err := database.GetSession(ck.Value)

	if err != nil || session.UserID == 0 {
		return user, err
	}

	user, err = database.GetUserByID(session.UserID)

	if err != nil || user.UserID == 0 {
		return user, err
	}

	ck.Path = "/"
	ck.MaxAge = 86400
	http.SetCookie(w, ck)

	return user, err
}
