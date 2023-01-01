package product

import (
	"time"

	"github.com/shariarfaisal/order-ms/brand"
	"gorm.io/gorm"
)

type ProductStatus string

const (
	ProductStatusPending ProductStatus = "pending"
	ProductStatusRejected ProductStatus = "rejected"
	ProductStatusApproved ProductStatus = "approved"
)

type ProductType string

const (
	ProductTypeSingle ProductType = "single"
	ProductTypeVariant ProductType = "variant"
)

type InventoryType string

const (
	InventoryTypePeriodic InventoryType = "periodic"
	InventoryTypePerpetual InventoryType = "perpetual"
)


type Product struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Type ProductType `json:"type"`
	Name        string `json:"name"`
	CategoryId int `json:"categoryId" gorm:"index"`
	Category brand.BrandCategory `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	Slug 	  string `json:"slug"`
	Images 	 []string `json:"image,omitempty" gorm:"type:jsonb"`
	Details    string `json:"details"`
	Price float32 `json:"price"`
	Status 	  ProductStatus `json:"status"`
	BrandId int `json:"brandId" gorm:"index"`
	Brand brand.Brand `json:"brand,omitempty" gorm:"foreignKey:BrandId"`
	IsAvailable bool `json:"isAvailable"`
	UseInventory bool `json:"useInventory"`
	InventoryType InventoryType `json:"inventoryType"`
	Stock int `json:"stock"`
	Variants []ProductVariant `json:"variants,omitempty" gorm:"foreignKey:ProductId"`
}

type ProductVariant struct {
	ID        uint `json:"id" gorm:"primarykey"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
	ProductId uint `json:"productId" gorm:"index"`
	Product Product `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	Title string `json:"title"`
	MinSelect int `json:"minSelect"`
	MaxSelect int `json:"maxSelect"`
	Items []VariantItem `json:"items,omitempty" gorm:"<-:create;foreignKey:VariantId"`
}

type VariantItem struct{
	ID uint `json:"id" gorm:"primarykey"`
	ProductId uint `json:"productId" gorm:"<-:create;index"`
	Product Product `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	VariantId uint `json:"variantId" gorm:"<-:create;index"`
	Variant ProductVariant `json:"variant,omitempty" gorm:"<-:create;foreignKey:VariantId"`
}

type PurchaseProduct struct {
	ID        uint `json:"id" gorm:"primarykey"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
	ProductId int `json:"productId" gorm:"index"`
	Product Product `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
	SellingPrice float64 `json:"sellingPrice"`
	ExpiredAt time.Time `json:"expiredAt" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
	BatchNumber string `json:"batchNumber"`
	Total float64 `json:"total"`
	PurchaseDate time.Time `json:"purchaseDate" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
}

type ProductDiscount struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductId int `json:"productId" gorm:"index"`
	Product Product `json:"product,omitempty" gorm:"<-:create;foreignKey:ProductId"`
	DiscountType string `json:"discountType"`
	DiscountValue float64 `json:"discountValue"`
	ValidFrom time.Time `json:"validFrom" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
	ValidTo time.Time `json:"validTo" gorm:"<-:create;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductVariant{})
	db.AutoMigrate(&PurchaseProduct{})
	db.AutoMigrate(&VariantItem{})
	db.AutoMigrate(&ProductDiscount{})
}