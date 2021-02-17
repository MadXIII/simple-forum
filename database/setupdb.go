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

	usersTable := `CREATE TABLE IF NOT EXISTS users (
		userid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		hash BLOB NOT NULL,
		salt TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	createUsersTable, err := db.Prepare(usersTable)
	errorhandle.CheckErr(err)
	_, err = createUsersTable.Exec()
	errorhandle.CheckErr(err)

	postsTable := `CREATE TABLE IF NOT EXISTS posts (
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
	);`

	createPostsTable, err := db.Prepare(postsTable)
	errorhandle.CheckErr(err)
	_, err = createPostsTable.Exec()
	errorhandle.CheckErr(err)

	categoryTable := `CREATE TABLE IF NOT EXISTS categories (
		categoryid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		categoryname TEXT NOT NULL
	);`
	createCategoryTable, err := db.Prepare(categoryTable)
	errorhandle.CheckErr(err)
	_, err = createCategoryTable.Exec()
	errorhandle.CheckErr(err)

	forumCategories := []string{
		"Sport",
		"Movies",
		"Music",
		"Other",
	}
	for _, category := range forumCategories {
		err = CreateCategory(category)
		if err != nil && err != ErrExists {
			errorhandle.CheckErr(err)
		}
	}
	postCategoriesTable := `CREATE TABLE IF NOT EXISTS postcategories (
		categoryid INTEGER NOT NULL,
		postid INTEGER NOT NULL,
		FOREIGN KEY (categoryid) REFERENCES categories(categoryid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`
	createPostCategoriesTable, err := db.Prepare(postCategoriesTable)
	errorhandle.CheckErr(err)
	_, err = createPostCategoriesTable.Exec()
	errorhandle.CheckErr(err)

	commentTable := `CREATE TABLE IF NOT EXISTS comments (
		commentid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postid INTEGER NOT NULL,
		username TEXT NOT NULL,
		text TEXT NOT NULL,
		imageexist  BYTE NOT NULL,
		date_time TIMESTAMP NOT NULL,
		timestring  TEXT NOT NULL,
		FOREIGN KEY (postid) REFERENCES posts(postid),
		FOREIGN KEY (username) REFERENCES users(username)
	);`

	createCommentTabel, err := db.Prepare(commentTable)
	errorhandle.CheckErr(err)
	_, err = createCommentTabel.Exec()
	errorhandle.CheckErr(err)

	sessionTable := `CREATE TABLE IF NOT EXISTS sessions (
		sessionid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userid INTEGER NOT NULL UNIQUE,
		timecreated TIMESTAMP NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid)
	);`

	createSessionTabel, err := db.Prepare(sessionTable)
	errorhandle.CheckErr(err)
	_, err = createSessionTabel.Exec()
	errorhandle.CheckErr(err)

	postLikesTable := `CREATE TABLE IF NOT EXISTS postlikes (
		userid INTEGER NOT NULL,
		postid INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (postid) REFERENCES posts(postid)
	);`
	createPostLikesTabel, err := db.Prepare(postLikesTable)
	errorhandle.CheckErr(err)
	_, err = createPostLikesTabel.Exec()
	errorhandle.CheckErr(err)

	commentLikesTable := `CREATE TABLE IF NOT EXISTS commentlikes (
		userid INTEGER NOT NULL,
		commentid INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (userid) REFERENCES users(userid),
		FOREIGN KEY (commentid) REFERENCES comments(commentid)
	);`

	createCommentLikesTable, err := db.Prepare(commentLikesTable)
	errorhandle.CheckErr(err)
	_, err = createCommentLikesTable.Exec()
	errorhandle.CheckErr(err)
}
