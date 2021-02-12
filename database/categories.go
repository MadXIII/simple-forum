package database

import "errors"

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
