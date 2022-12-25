package rider

import "gorm.io/gorm"

type Rider struct {
	gorm.Model
	Name   string
	Phone  string
	Image  string
	Type   string
	Status string
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Rider{})
}