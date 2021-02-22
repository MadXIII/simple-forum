package database

import (
	"database/sql"
	"forum/errorhandle"

	_ "github.com/mattn/go-sqlite3"
)

var err error
var db *sql.DB

// SetUp - ...
func SetUp() {
	db, err = sql.Open("sqlite3", "./database.db")
	errorhandle.CheckErr(err)

	usersTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
		userid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		hash BLOB NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`)
	errorhandle.CheckErr(err)

	_, err = usersTable.Exec()
	errorhandle.CheckErr(err)

	postsTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS posts (
		postid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userid INTEGER NOT NULL,
		username TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		imageexist  BYTE NOT NULL,
		date_time TIMESTAMP NOT NULL,
		timestring  TEXT NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (username) REFERENCES users(username)
	);`)
	errorhandle.CheckErr(err)

	_, err = postsTable.Exec()
	errorhandle.CheckErr(err)

	categoryTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS categories (
		categoryid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		categoryname TEXT NOT NULL UNIQUE
	);`)
	errorhandle.CheckErr(err)

	_, err = categoryTable.Exec()
	errorhandle.CheckErr(err)

	postCategories := []string{
		"Club",
		"Matches",
		"Team",
		"Transfers",
	}
	for _, category := range postCategories {
		err = CreateCategory(category)
		if err != nil && err != ErrExists {
			errorhandle.CheckErr(err)
		}
	}
	postCategoriesTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS postcategories (
		categoryid INTEGER NOT NULL,
		postid INTEGER NOT NULL,
		FOREIGN KEY (categoryid) REFERENCES categories(categoryid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`)
	errorhandle.CheckErr(err)

	_, err = postCategoriesTable.Exec()
	errorhandle.CheckErr(err)

	commentTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS comments (
		commentid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postid INTEGER NOT NULL,
		username TEXT NOT NULL,
		text TEXT NOT NULL,
		date_time TIMESTAMP NOT NULL,
		timestring  TEXT NOT NULL,
		FOREIGN KEY (postid) REFERENCES posts(postid),
		FOREIGN KEY (username) REFERENCES users(username)
	);`)
	errorhandle.CheckErr(err)
	_, err = commentTable.Exec()
	errorhandle.CheckErr(err)

	sessionTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS sessions (
		sessionid STRING NOT NULL PRIMARY KEY,
		userid INTEGER NOT NULL UNIQUE,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`)
	errorhandle.CheckErr(err)

	_, err = sessionTable.Exec()
	errorhandle.CheckErr(err)

	postLikesTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS postlikes (
		userid INTEGER NOT NULL,
		postid INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`)
	errorhandle.CheckErr(err)

	_, err = postLikesTable.Exec()
	errorhandle.CheckErr(err)

	commentLikesTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS commentlikes (
		userid INTEGER NOT NULL,
		commentid INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (commentid) REFERENCES comments(commentid)
	);`)
	errorhandle.CheckErr(err)

	_, err = commentLikesTable.Exec()
	errorhandle.CheckErr(err)
}
