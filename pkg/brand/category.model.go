package brand

import (
	"github.com/shariarfaisal/order-ms/pkg/category"
)

type BrandCategory struct {
	ID         uint                      `json:"id,omitempty" gorm:"primarykey"`
	BrandId    uint                      `json:"brandId,omitempty" gorm:"index"`
	Brand      *Brand                    `json:"brand,omitempty" gorm:"foreignKey:BrandId"`
	CategoryId uint                      `json:"categoryId,omitempty" gorm:"index"`
	Category   *category.ProductCategory `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	IsActive   bool                      `json:"isActive,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Slug       string                    `json:"slug,omitempty"`
}
