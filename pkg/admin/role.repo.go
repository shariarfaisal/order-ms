package admin

import "gorm.io/gorm"

type RoleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{DB: db}
}

func (r *RoleRepo) GetByName(name string) (role Role, err error) {
	err = r.DB.Where(" name = ?", name).First(&role).Error
	return
}
