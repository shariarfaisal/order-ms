package brand

import (
	"github.com/shariarfaisal/order-ms/category"
)

type BrandCategory struct {
	ID 	  uint `json:"id" gorm:"primarykey"`
	BrandId     uint                     `json:"brandId" gorm:"index"`
	Brand       Brand                   `json:"brand" gorm:"foreignKey:BrandId"`
	CategoryId  uint                     `json:"categoryId" gorm:"index"`
	Category    category.ProductCategory `json:"category" gorm:"foreignKey:CategoryId"`
	IsActive    bool                    `json:"isActive"`
	Name 	  string                  `json:"name"`
	Slug 	  string                  `json:"slug"`
}