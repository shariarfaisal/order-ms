package customer

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string
	Image    string
	Phone    string
	Username string
	Email    string
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
}