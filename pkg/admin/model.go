package admin

import "gorm.io/gorm"

type Admin struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Admin{}, &Role{})
}
