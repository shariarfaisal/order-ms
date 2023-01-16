package brand

import "gorm.io/gorm"

type PartnerRepo struct {
	DB *gorm.DB
}

func NewPartnerRepo(db *gorm.DB) *PartnerRepo {
	return &PartnerRepo{DB: db}
}

func (r *PartnerRepo) GetById(id uint) (partner Partner, err error) {
	err = r.DB.First(&partner, id).Error
	return partner, err
}

func (r *PartnerRepo) GetItems() (partners []Partner, err error) {
	err = r.DB.Find(&partners).Error
	return partners, err
}
