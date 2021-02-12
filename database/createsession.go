package database

import (
	"forum/models"
)

func CreateSession(newSession models.Session) error {
	createSession, err := db.Prepare(`
		REPLACE INTO sessions
		(sessionid, userid, timecreated)
		VALUES (?, ?, ?);
	`)
	if err != nil {
		return err
	}
	_, err = createSession.Exec(
		newSession.SessionID,
		newSession.UserID,
		newSession.TimeCreated,
	)
	if err != nil {
		return err
	}
	return err
}
