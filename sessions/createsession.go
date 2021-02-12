package sessions

import (
	"forum/database"
	"forum/models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreateSession(uid int, w http.ResponseWriter) error {
	sid := uuid.NewV4().String()
	cook := &http.Cookie{
		Name:   "session",
		Value:  sid,
		MaxAge: 86400,
		Path:   "/",
	}
	session := models.Session{
		SessionID:   cook.Value,
		UserID:      uid,
		TimeCreated: time.Now(),
	}
	err := database.CreateSession(session)
	if err != nil {
		return err
	}
	http.SetCookie(w, cook)
	return nil
}
