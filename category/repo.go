package category

// getCategories
func GetCategories() ([]ProductCategory, error) {
	categories := []ProductCategory{}
	db.Find(&categories)
	return categories, nil
}

// createCategory
func CreateCategory(category ProductCategory) (ProductCategory, error) {
	err := db.Create(&category).Error
	return category, err
}


// getCategoryById
func GetCategoryById(id int) (ProductCategory, error) {
	var category ProductCategory
	err := db.Where("id = ?", id).First(&category).Error
	return category, err
}

// updateCategory
func UpdateCategoryById(id int, category ProductCategory) (ProductCategory, error) {
	err := db.Model(&ProductCategory{}).Where("id = ?", id).Updates(category).Error
	return category, err
}


// deleteCategory
func DeleteCategoryById(id int) error {
	err := db.Delete(&ProductCategory{}, id).Error
	return err
}