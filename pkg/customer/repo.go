package customer

import (
	"errors"

	"gorm.io/gorm"
)

func GetAddressById(id int) (CustomerAddress, error) {
	var address CustomerAddress
	err := db.Where("id = ?", id).First(&address).Error
	return address, err
}

func GetCustomerById(id int) (Customer, error) {
	var customer Customer
	err := db.Where("id = ?", id).First(&customer).Error
	return customer, err
}

func GetCustomerByEmail(email string) (Customer, error) {
	var customer Customer
	err := db.Where("email = ?", email).First(&customer).Error
	return customer, err
}

func GetCustomerByPhone(phone string) (Customer, error) {
	var customer Customer
	err := db.Where("phone = ?", phone).First(&customer).Error
	return customer, err
}

func GetBy(key string, value interface{}) (Customer, error) {
	var customer Customer
	err := db.Where(key+" = ?", value).First(&customer).Error
	return customer, err
}

func IsCustomerExist(key string, value string) (bool, error) {
	var customer Customer
	err := db.Where(key+" = ?", value).First(&customer).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}
