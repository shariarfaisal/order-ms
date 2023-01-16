package category

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint      `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Name      string    `json:"name,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	Icon      string    `json:"icon,omitempty"`
	Image     string    `json:"image,omitempty"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&ProductCategory{})
}
