package product

import (
	"github.com/shariarfaisal/order-ms/brand"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"<-:create"`
	Slug 	  string `json:"slug" gorm:"<-:create"`
	Details    string `json:"details" gorm:"<-:create"`
	Price 	int    `json:"price" gorm:"<-:create"`
	BrandId int `json:"brand_id" gorm:"<-:create" gorm:"index"`
	Brand   brand.Brand `json:"brand" gorm:"<-:create"`
	CategoryId int `json:"category_id" gorm:"<-:create" gorm:"index"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}