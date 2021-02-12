package database

import (
	"forum/models"
)

func GetUserByID(uid int) (models.User, error) {
	var user models.User
	row, err := db.Query(`
		SELECT * 
		FROM users
		WHERE userid = ?
	`, uid)
	defer row.Close()

	if err != nil {
		return user, err
	}
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email)
	}
	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	row, err := db.Query(`
		SELECT * 
		FROM users
		WHERE userid = ?
	`, username)
	defer row.Close()

	if err != nil {
		return user, err
	}
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email)
	}
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row, err := db.Query(`
		SELECT * 
		FROM users
		WHERE userid = ?
	`, email)
	defer row.Close()

	if err != nil {
		return user, err
	}
	for row.Next() {
		row.Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email)
	}
	return user, err
}