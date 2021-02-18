package database

import (
	"errors"
	"forum/models"
)

var ErrExists error = errors.New("Already exists")

//CreateCategory - ...
func CreateCategory(category string) error {
	createCategory, err := db.Prepare(`
	INSERT INTO categories
	(categoryname) 
	VALUES (?);
	`)
	if err != nil {
		return err
	}
	_, err = createCategory.Exec(category)
	if err != nil && err.Error() == "UNIQUE constraint failed: categories.categoryname" {
		return ErrExists
	}
	return err
}

func GetCategories() ([]string, error) {
	var categories []string

	row, err := db.Query(`
		SELECT categoryname
		FROM categories
	`)
	defer row.Close()

	if err != nil {
		return categories, err
	}

	for row.Next() {
		var category string
		row.Scan(&category)
		categories = append(categories, category)
	}
	return categories, err
}

func insertPostIntoCategories(postId int, categories []string) error {
	insert, err := db.Prepare(`
		INSERT INTO postcategories
		(categoryid, postid)
		SELECT categoryid, ?
		FROM categories
		WHERE categoryname = ?
	`)
	if err != nil {
		return err
	}
	for _, category := range categories {
		_, err = insert.Exec(postId, category)
		if err != nil {
			return err
		}
	}
	return err
}

func getPostCategories(post *models.Post) error {
	var categories []string

	row, err := db.Query(`
		SELECT categories.categoryname
		FROM postcategories
		INNER JOIN categories
		ON postcategories.categoryid = categories.categoryid
		WHERE postcategories.postid = ?
	`, post.PostID)
	defer row.Close()

	if err != nil {
		return err
	}

	for row.Next() {
		var category string
		row.Scan(&category)
		categories = append(categories, category)
	}
	post.Categories = categories

	return err
}
