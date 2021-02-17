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

func GetSession(sessionID string) (models.Session, error) {
	var session models.Session

	if err := db.QueryRow(`
		SELECT *
		FROM sessions
		WHERE sessionid = ?
	`, sessionID).Scan(&session.SessionID, &session.UserID, &session.TimeCreated); err != nil {
		return session, err
	}

	return session, err
}
