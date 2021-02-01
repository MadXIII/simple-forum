package database

//CreateCategory - ...
func CreateCategory(category string) error {
	createCategory, err := db.Prepare(`INSERT INTO categories(categoryname) VALUES (?);`)
	if err != nil {
		return err
	}
	_, err = createCategory.Exec(category)
	if err != nil {
		return err
	}
	return err
}
