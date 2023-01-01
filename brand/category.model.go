package brand

import (
	"github.com/shariarfaisal/order-ms/category"
	"gorm.io/gorm"
)

type BrandCategory struct {
	gorm.Model
	BrandId     int                     `json:"brandId" gorm:"index"`
	Brand       Brand                   `json:"brand" gorm:"foreignKey:BrandId"`
	CategoryId  int                     `json:"categoryId" gorm:"index"`
	Category    category.ProductCategory `json:"category" gorm:"foreignKey:CategoryId"`
	IsActive    bool                    `json:"isActive"`
	Name 	  string                  `json:"name"`
	Slug 	  string                  `json:"slug"`
}