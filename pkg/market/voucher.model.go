package market

import (
	"time"

	"github.com/shariarfaisal/order-ms/pkg/utils"
	"gorm.io/gorm"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
	DiscountTypeFixed      DiscountType = "fixed"
)

type ApplyOn string

const (
	ApplyOnProduct        ApplyOn = "product"
	ApplyOnCategory       ApplyOn = "category"
	ApplyOnRestaurant     ApplyOn = "restaurant"
	ApplyOnDeliveryCharge ApplyOn = "delivery_charge"
)

// scope
// audience
// usage

type Voucher struct {
	ID                    uint           `json:"id"`
	CreatedAt             time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt             time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title                 string         `json:"title"`
	SubTitle              string         `json:"sub_title"`
	Code                  string         `json:"code"`
	Auth                  string         `json:"auth"`
	Discount              float64        `json:"discount"`
	DiscountType          DiscountType   `json:"discount_type"`
	MinOrderAmount        float64        `json:"min_order_amount"`
	MaxDiscountAmount     float64        `json:"max_discount_amount"`
	StartDate             time.Time      `json:"start_date" gorm:"type:timestamptz;"`
	EndDate               time.Time      `json:"end_date" gorm:"type:timestamptz;"`
	Hours                 []int          `json:"hours" gorm:"type:integer[]"`
	Days                  []int          `json:"days" gorm:"type:integer[]"`
	IsActive              bool           `json:"is_active"`
	Platform              utils.Platform `json:"platform"`
	ApplyOn               ApplyOn        `json:"apply_on"` // product, category, restaurant, delivery charge
	RestrictedRestaurants []uint         `json:"restricted_restaurants" gorm:"type:integer[]"`
	RestrictedCategories  []uint         `json:"restricted_categories" gorm:"type:integer[]"`
	Categories            []uint         `json:"categories" gorm:"type:integer[]"`
	MaxUsage              int            `json:"max_usage"`
	MaxUsagePerUser       int            `json:"max_usage_per_user"`
	TotalUsed             int            `json:"total_used"`
	EligibleUsers         []uint         `json:"eligible_users" gorm:"type:integer[]"`
	UserMinOrderCount     int            `json:"user_min_order_count"`
	UserMaxOrderCount     int            `json:"user_max_order_count"`
	GeoJson               string         `json:"geo_json"`
}
