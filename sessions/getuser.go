package sessions

import (
	"forum/database"
	"forum/models"
	"net/http"
)

func GetUser(w http.ResponseWriter, r http.Request) (models.User, error) {
	var user models.User

	cook, err := r.Cookie("session")
	if err != nil {
		return user, err
	}
	session, err := database.GetSession(cook.Value)

	if err != nil || session.UserID == 0 {
		return user, err
	}

	user, err 
}
