package category

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint `json:"id" gorm:"primarykey"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Icon  string `json:"icon"`
	Image string `json:"image"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&ProductCategory{})
}