package market

import (
	"errors"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	DB *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{DB: db}
}

func (r *CustomerRepo) GetAddressById(id int) (CustomerAddress, error) {
	var address CustomerAddress
	err := r.DB.Where("id = ?", id).First(&address).Error
	return address, err
}

func (r *CustomerRepo) GetById(id int) (Customer, error) {
	var customer Customer
	err := r.DB.Where("id = ?", id).First(&customer).Error
	return customer, err
}

func (r *CustomerRepo) GetByEmail(email string) (Customer, error) {
	var customer Customer
	err := r.DB.Where("email = ?", email).First(&customer).Error
	return customer, err
}

func (r *CustomerRepo) GetByPhone(phone string) (Customer, error) {
	var customer Customer
	err := r.DB.Where("phone = ?", phone).First(&customer).Error
	return customer, err
}

func (r *CustomerRepo) GetBy(key string, value interface{}) (Customer, error) {
	var customer Customer
	err := r.DB.Where(key+" = ?", value).First(&customer).Error
	return customer, err
}

func (r *CustomerRepo) IsExist(key string, value string) (bool, error) {
	var customer Customer
	err := r.DB.Where(key+" = ?", value).First(&customer).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}
