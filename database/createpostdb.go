package database

import (
	"forum/models"
)

func CreatePost(newPost *models.Post) error {
	createPost, err := db.Prepare(`
		INSERT INTO posts
		(userid, username, title, content, imageexist, date_time, timestring)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}

	res, err := createPost.Exec(
		newPost.UserID,
		newPost.Username,
		newPost.Title,
		newPost.Content,
		newPost.ImageExist,
		newPost.DateTime,
		newPost.TimeString,
	)
	if err != nil {
		return err
	}

	PostID, err := res.LastInsertId()
	newPost.PostID = int(PostID)
	if err != nil {
		return err
	}

	err = insertPostIntoCategories(newPost.PostID, newPost.Categories)
	if err != nil {
		return err
	}

	return err
}
