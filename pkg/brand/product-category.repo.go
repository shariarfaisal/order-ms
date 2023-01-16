package brand

import "gorm.io/gorm"

type ProductCategoryRepo struct {
	DB *gorm.DB
}

func NewProductCategoryRepo(db *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{DB: db}
}

// getCategories
func (r *ProductCategoryRepo) GetItems() ([]ProductCategory, error) {
	categories := []ProductCategory{}
	r.DB.Find(&categories)
	return categories, nil
}

// getCategoryById
func (r *ProductCategoryRepo) GetById(id uint) (ProductCategory, error) {
	var category ProductCategory
	err := r.DB.Where("id = ?", id).First(&category).Error
	return category, err
}

// updateCategory
func (r *ProductCategoryRepo) UpdateById(id uint, category ProductCategory) (ProductCategory, error) {
	err := r.DB.Model(&ProductCategory{}).Where("id = ?", id).Updates(category).Error
	return category, err
}

// deleteCategory
func (r *ProductCategoryRepo) DeleteById(id uint) error {
	err := r.DB.Delete(&ProductCategory{}, id).Error
	return err
}

func (r *ProductCategoryRepo) IsExists(key string, value interface{}) (bool, error) {
	var count int64
	err := r.DB.Model(&ProductCategory{}).Where(key+" = ?", value).Count(&count).Error
	return count > 0, err
}
