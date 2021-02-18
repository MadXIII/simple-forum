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

	row, err := db.Query(`
		SELECT *
		FROM sessions
		WHERE sessionid = ?
	`, sessionID)
	defer row.Close()

	if err != nil {
		return session, err
	}
	for row.Next() {
		row.Scan(&session.SessionID, &session.UserID, &session.TimeCreated)
	}

	return session, err
}

//QueryRow
