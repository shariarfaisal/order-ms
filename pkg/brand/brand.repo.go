package brand

import "gorm.io/gorm"

type BrandRepo struct {
	DB *gorm.DB
}

func NewBrandRepo(db *gorm.DB) *BrandRepo {
	return &BrandRepo{DB: db}
}

func (r *BrandRepo) GetItems() (brands []Brand, err error) {
	err = r.DB.Find(&brands).Error
	return brands, err
}

func (r *BrandRepo) GetById(id uint) (brand Brand, err error) {
	err = r.DB.First(&brand, id).Error
	return brand, err
}

func (r *BrandRepo) GetByNameAndPartnerId(name string, partnerId uint) (brand Brand, err error) {
	err = r.DB.Where("name = ? AND partner_id = ?", name, partnerId).First(&brand).Error
	return brand, err
}
