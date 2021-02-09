package database

//IsUsernameExists - ...
func IsUsernameExists(username string) bool {
	row, err := db.Query(`
	SELECT 1
	FROM users
	WHERE username = ?
	`, username)
	defer row.Close()
	if err != nil {
		return false
	}
	return row.Next()
}

//IsEmailExists - ...
func IsEmailExists(email string) bool {
	row, err := db.Query(`
	SELECT 1
	FROM users
	WHERE email = ?
	`, email)
	defer row.Close()
	if err != nil {
		return false
	}
	return row.Next()
}
