package brand

import (
	"time"
)

type ProductCategory struct {
	ID        uint            `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time       `json:"createdAt,omitempty"`
	UpdatedAt time.Time       `json:"updatedAt,omitempty"`
	Name      string          `json:"name,omitempty"`
	Slug      string          `json:"slug,omitempty"`
	Icon      string          `json:"icon,omitempty"`
	Image     string          `json:"image,omitempty"`
	IsActive  bool            `json:"isActive,omitempty"`
	Items     []BrandCategory `json:"items,omitempty" gorm:"foreignKey:CategoryId"`
}

type BrandCategory struct {
	ID         uint            `json:"id,omitempty" gorm:"primarykey"`
	BrandId    uint            `json:"brandId,omitempty" gorm:"index"`
	Brand      Brand           `json:"brand,omitempty" gorm:"foreignKey:BrandId"`
	CategoryId uint            `json:"categoryId,omitempty" gorm:"index"`
	Category   ProductCategory `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	IsActive   bool            `json:"isActive,omitempty"`
	Name       string          `json:"name,omitempty"`
	Slug       string          `json:"slug,omitempty"`
}
