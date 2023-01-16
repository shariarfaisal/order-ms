package brand

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProductStatus string

const (
	ProductStatusPending  ProductStatus = "pending"
	ProductStatusRejected ProductStatus = "rejected"
	ProductStatusApproved ProductStatus = "approved"
)

type ProductType string

const (
	ProductTypeSingle  ProductType = "single"
	ProductTypeVariant ProductType = "variant"
)

type InventoryType string

const (
	InventoryTypePeriodic  InventoryType = "periodic"
	InventoryTypePerpetual InventoryType = "perpetual"
)

type Product struct {
	gorm.Model
	ID            uint             `json:"id" gorm:"primarykey"`
	Type          ProductType      `json:"type"`
	Name          string           `json:"name"`
	CategoryId    uint             `json:"categoryId" gorm:"index"`
	Category      BrandCategory    `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	Slug          string           `json:"slug"`
	Images        pq.StringArray   `json:"image,omitempty" gorm:"type:text[]"`
	Details       string           `json:"details"`
	Price         float32          `json:"price"`
	Status        ProductStatus    `json:"status"`
	BrandId       uint             `json:"brandId" gorm:"index"`
	Brand         Brand            `json:"brand,omitempty" gorm:"foreignKey:BrandId"`
	IsAvailable   bool             `json:"isAvailable"`
	UseInventory  bool             `json:"useInventory"`
	InventoryType InventoryType    `json:"inventoryType"`
	Stock         int              `json:"stock"`
	Variants      []ProductVariant `json:"variants,omitempty" gorm:"foreignKey:ProductId"`
	VariantId     uint             `json:"variantId" gorm:"index"`
	Variant       *ProductVariant  `json:"variant,omitempty" gorm:"foreignKey:VariantId"`
}

type ProductVariant struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductId uint      `json:"productId" gorm:"index"`
	Product   Product   `json:"-" gorm:"<-:create;foreignKey:ProductId"`
	Title     string    `json:"title"`
	MinSelect int       `json:"minSelect"`
	MaxSelect int       `json:"maxSelect"`
	Items     []Product `json:"items,omitempty" gorm:"many2many:product_variant_items;"`
}

type PurchaseProduct struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ProductId    int       `json:"productId" gorm:"index"`
	Product      Product   `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	SellingPrice float64   `json:"sellingPrice"`
	ExpiredAt    time.Time `json:"expiredAt" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
	BatchNumber  string    `json:"batchNumber"`
	Total        float64   `json:"total"`
	PurchaseDate time.Time `json:"purchaseDate" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
}

type ProductDiscount struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	ProductId     int       `json:"productId" gorm:"index"`
	Product       Product   `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	DiscountType  string    `json:"discountType"`
	DiscountValue float64   `json:"discountValue"`
	ValidFrom     time.Time `json:"validFrom" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
	ValidTo       time.Time `json:"validTo" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
}
