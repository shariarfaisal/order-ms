package admin

import (
	"errors"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func (r *AdminRepo) GetByEmail(email string) (admin Admin, err error) {
	err = r.DB.Where(" email = ?", email).First(&admin).Error
	return
}

func (r *AdminRepo) GetBy(key string, value interface{}) (admin Admin, err error) {
	err = r.DB.Where(key+" = ?", value).First(&admin).Error
	return
}

func (r *AdminRepo) IsExist(key string, value string) (exists bool, err error) {
	result := Admin{}
	err = r.DB.Model(&Admin{}).Select(key).Where(key+" = ?", value).First(&result).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err == nil {
		return true, err
	}
	return
}
