package brand

import "gorm.io/gorm"

type BrandCategoryRepo struct {
	DB *gorm.DB
}

func NewBrandCategoryRepo(db *gorm.DB) *BrandCategoryRepo {
	return &BrandCategoryRepo{DB: db}
}

func (r *BrandCategoryRepo) GetById(id uint) (category BrandCategory, err error) {
	err = r.DB.First(&category, id).Error
	return category, err
}

func (r *BrandCategoryRepo) GetByPCategoryId(id uint) (category BrandCategory, err error) {
	err = r.DB.Where("category_id = ?", id).First(&category).Error
	return category, err
}

func (r *BrandCategoryRepo) GetItems() (categories []BrandCategory, err error) {
	err = r.DB.Find(&categories).Error
	return categories, err
}
