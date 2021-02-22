package database

import "forum/models"

func CreateUser(newUser *models.User) error {
	createUser, err := db.Prepare(`
		INSERT INTO users
		(username, hash, email)
		VALUES (?, ?, ?);
	`)
	if err != nil {
		return err
	}

	res, err := createUser.Exec(
		newUser.Username,
		newUser.Hash,
		newUser.Email,
	)
	if err != nil {
		return err
	}
	userid, err := res.LastInsertId()
	newUser.UserID = int(userid)
	if err != nil {
		return err
	}
	return err
}
