package database

import (
	"forum/models"
)

func GetUserByID(uid int) (models.User, error) {
	var user models.User
	if err := db.QueryRow(`
		SELECT * 
		FROM users
		WHERE userid = ?
	`, uid).Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email); err != nil {
		return user, err
	}
	return user, err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := db.QueryRow(`
		SELECT * 
		FROM users
		WHERE username = ?
	`, username).Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email); err != nil {
		return user, err
	}
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := db.QueryRow(`
		SELECT * 
		FROM users
		WHERE email = ?
	`, email).Scan(&user.UserID, &user.Username, &user.Hash, &user.Salt, &user.Email); err != nil {
		return user, err
	}
	return user, err
}
