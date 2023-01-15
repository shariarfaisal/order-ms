package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)
}

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
