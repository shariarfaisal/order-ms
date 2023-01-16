package admin

import (
	"errors"

	"gorm.io/gorm"
)

func GetRoleByName(name string) (role Role, err error) {
	err = db.Where(" name = ?", name).First(&role).Error
	return
}

func GetAdminByEmail(email string) (admin Admin, err error) {
	err = db.Where(" email = ?", email).First(&admin).Error
	return
}

func GetBy(key string, value interface{}) (admin Admin, err error) {
	err = db.Where(key+" = ?", value).First(&admin).Error
	return
}

func IsAdminExist(key string, value string) (exists bool, err error) {
	result := Admin{}
	err = db.Model(&Admin{}).Select(key).Where(key+" = ?", value).First(&result).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err == nil {
		return true, err
	}
	return
}
