package order

import (
	"time"

	"github.com/shariarfaisal/order-ms/pkg/admin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"github.com/shariarfaisal/order-ms/pkg/market"
	"github.com/shariarfaisal/order-ms/pkg/rider"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPending    OrderStatus = "pending"
	OrderAccepted   OrderStatus = "accepted"
	OrderPreparing  OrderStatus = "preparing"
	OrderReady      OrderStatus = "ready"
	OrderDispatched OrderStatus = "dispatched"
	OrderDelivered  OrderStatus = "delivered"
	OrderCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID              uint             `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"`    // primary key
	CreatedAt       time.Time        `json:"createdAt" gorm:"<-:create" gorm:"type:timestamp;"` // created at
	UpdatedAt       time.Time        `json:"updatedAt" gorm:"<-:create" gorm:"type:timestamp"`  // updated at
	DeletedAt       gorm.DeletedAt   `json:"deletedAt" gorm:"<-:create" gorm:"type:timestamp"`  // deleted at
	DeliveredTo     uint             `json:"deliveredTo" gorm:"<-:create"`
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress" gorm:"foreignKey:DeliveredTo;references:ID"`
	Status          OrderStatus      `json:"status" gorm:"<-:create"`
	Platform        utils.Platform   `json:"platform" gorm:"<-:create"`
	DispatchTime    string           `json:"dispatchedTime" gorm:"<-:create"`
	RiderNote       string           `json:"riderNote" gorm:"<-:create"`
	ConfirmedAt     time.Time        `json:"confirmedAt" gorm:"<-:create" gorm:"type:timestamp"`
	AssignedTo      uint             `json:"assignedTo" gorm:"<-:create" gorm:"index"`
	AssignedRider   *AssignedRider   `json:"assignedRider" gorm:"foreignKey:AssignedTo;references:ID"`
	HubId           uint             `json:"hubId" gorm:"<-:create" gorm:"index"`
	Hub             hub.Hub          `json:"hub" gorm:"foreignKey:HubId"`
	EDT             int              `json:"edt" gorm:"<-:create"`
	PaymentMethod   string           `json:"paymentMethod" gorm:"<-:create"`
	PaymentStatus   string           `json:"paymentStatus" gorm:"<-:create"`
	ChargesId       uint             `json:"chargesId" gorm:"<-:create" gorm:"index"`
	Charges         *OrderCharges    `json:"charges" gorm:"foreignKey:ChargesId;references:ID"`
	Pickups         []*Pickup        `json:"pickups" gorm:"foreignKey:OrderId;references:ID"`
	Items           []*OrderItem     `json:"items" gorm:"foreignKey:OrderId;references:ID"`
	Timeline        []*OrderTimeline `json:"timeline" gorm:"foreignKey:OrderId;references:ID"`
	CompletedAt     time.Time        `json:"completedAt" gorm:"<-:create" gorm:"type:timestamp"`
}

type OrderTimeline struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	OrderId     uint      `json:"order_id" gorm:"<-:create" gorm:"index"`
	Order       Order     `json:"order" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // OrderId is the foreign key
	OrderNumber string    `json:"order_number" gorm:"<-:create"`
	PickupId    uint      `json:"pickup_id" gorm:"<-:create" gorm:"index"`
	Pickup      *Pickup   `json:"pickup" gorm:"foreignKey:PickupId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // PickupId is the foreign key
	Type        string    `json:"type" gorm:"<-:create"`
	Note        string    `json:"note" gorm:"<-:create"`
	CreatedAt   time.Time `json:"createdAt" gorm:"<-:create" gorm:"type:timestamp;"` // created at
}

type OrderCharges struct {
	ID              uint    `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	Total           float64 `json:"total"`
	ProductDiscount float64 `json:"productDiscount"`
	VoucherDiscount float64 `json:"voucherDiscount"`
	Voucher         struct {
		Code           string  `json:"code"`
		DiscountAmount float64 `json:"discountAmount"`
		DiscountType   string  `json:"discountType"`
	} `json:"voucher" gorm:"type:jsonb"`
	DeliveryCharge   float64 `json:"deliveryCharge"`
	DeliveryDiscount float64 `json:"deliveryDiscount"`
	TotalDiscount    float64 `json:"totalDiscount"`
	ServiceCharge    float64 `json:"serviceCharge"`
	SubTotal         float64 `json:"subTotal"`
}

type PickupStatus string

const (
	PickupAccepted PickupStatus = "accepted"
	PickupRejected PickupStatus = "rejected"
)

type Pickup struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"`    // primary key
	CreatedAt   time.Time      `json:"createdAt" gorm:"<-:create" gorm:"type:timestamp;"` // created at
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"<-:create" gorm:"type:timestamp"`  // updated at
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"<-:create" gorm:"type:timestamp"`  // deleted at
	BrandId     uint           `json:"brand_id" gorm:"<-:create" gorm:"index"`
	Brand       brand.Brand    `json:"brand" gorm:"foreignKey:BrandId"`
	OrderId     uint           `json:"order_id" gorm:"<-:create" gorm:"index"`
	Order       Order          `json:"order" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // OrderId is the foreign key
	OrderNumber string         `json:"order_number" gorm:"<-:create"`
	Note        string         `json:"note" gorm:"<-:create"`
	Total       float64        `json:"total" gorm:"<-:create"`
	Status      PickupStatus   `json:"status" gorm:"<-:create"`
	Items       []*OrderItem   `json:"items" gorm:"foreignKey:PickupId;references:ID"`
}

type OrderItem struct {
	gorm.Model
	ProductId uint           `json:"product_id" gorm:"<-:create" gorm:"index"`
	Product   *brand.Product `json:"product" gorm:"<-:create;foreignKey:ProductId;references:ID"`
	Quantity  int            `json:"quantity" gorm:"<-:create"`
	SaleUnit  float64        `json:"sale_unit" gorm:"<-:create"`
	Total     float64        `json:"total" gorm:"<-:create"`
	Discount  float64        `json:"discount" gorm:"<-:create"`
	PickupId  uint           `json:"pickup_id" gorm:"<-:create" gorm:"index"`
	Pickup    *Pickup        `json:"pickup" gorm:"foreignKey:PickupId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderId   uint           `json:"order_id" gorm:"<-:create" gorm:"index"`
	Order     *Order         `json:"order" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DeliveryAddress struct {
	gorm.Model
	Name       string           `json:"name" gorm:"<-:create"`
	Phone      string           `json:"phone" gorm:"<-:create" gorm:"index"`
	Address    string           `json:"address" gorm:"<-:create"`
	Area       string           `json:"area" gorm:"<-:create"`
	Lat        float64          `json:"lat" gorm:"<-:create"`
	Lng        float64          `json:"lng" gorm:"<-:create"`
	CustomerId uint             `json:"customer_id" gorm:"<-:create" gorm:"index"`
	Customer   *market.Customer `json:"customer" gorm:"<-:create;foreignKey:CustomerId;references:ID"`
}

type AssignedRider struct {
	gorm.Model
	RiderId uint
	Rider   *rider.Rider `json:"rider" gorm:"<-:create;foreignKey:RiderId;references:ID"`
	Note    string       `json:"note" gorm:"<-:create"`
}

type CartItem struct {
	ID         uint             `json:"id"`
	Quantity   uint             `json:"quantity"`
	CustomerId uint             `json:"user_id"`
	Customer   *market.Customer `json:"customer" gorm:"<-:create;foreignKey:CustomerId;references:ID"`
	ProductId  uint             `json:"product_id"`
	Product    *brand.Product   `json:"product" gorm:"<-:create;foreignKey:ProductId;references:ID"`
	CreatedAt  time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt   `json:"deleted_at"` // soft delete
}

type OrderIssue struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	CreatedAt time.Time      `json:"createdAt" gorm:"type:timestamp;"`               // created at
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:timestamp"`                // updated at
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp"`                // deleted at
	OrderId   uint           `json:"order_id" gorm:"index"`
	Order     *Order         `json:"order" gorm:"foreignKey:OrderId;references:ID"`
	IssueType string         `json:"issue_type"`
	Note      string         `json:"note"`
	RaisedBy  string         `json:"raised_by"`
	Status    string         `json:"status"`
	IssuedBy  string         `json:"issued_by"`
	IsRefund  bool           `json:"is_refund"`
	Refunds   []*OrderRefund `json:"refunds" gorm:"foreignKey:IssueId;references:ID"`
}

type RestaurantPenalty struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	CreatedAt time.Time      `json:"createdAt" gorm:"type:timestamp;"`               // created at
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:timestamp"`                // updated at
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp"`                // deleted at
	OrderId   uint           `json:"order_id" gorm:"index"`
	Order     *Order         `json:"order" gorm:"foreignKey:OrderId;references:ID"`
	Amount    float64        `json:"amount"`
	Note      string         `json:"note"`
}

type RiderPenalty struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	CreatedAt time.Time      `json:"createdAt" gorm:"type:timestamp;"`               // created at
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:timestamp"`                // updated at
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp"`                // deleted at
	OrderId   uint           `json:"order_id" gorm:"index"`
	Order     *Order         `json:"order" gorm:"foreignKey:OrderId;references:ID"`
	Amount    float64        `json:"amount"`
	Note      string         `json:"note"`
}

type OrderRefund struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement;uniqueIndex"` // primary key
	CreatedAt    time.Time      `json:"createdAt" gorm:"type:timestamp;"`               // created at
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"type:timestamp"`                // updated at
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"type:timestamp"`                // deleted at
	OrderId      uint           `json:"order_id" gorm:"index"`
	Order        *Order         `json:"order" gorm:"foreignKey:OrderId;references:ID"`
	Amount       float64        `json:"amount"`
	Note         string         `json:"note"`
	IssueId      uint           `json:"issue_id"`
	Issue        *OrderIssue    `json:"issue" gorm:"foreignKey:IssueId;references:ID"`
	RefundedById uint           `json:"refunded_by_id"`
	RefundedBy   *admin.Admin   `json:"refunded" gorm:"foreignKey:RefundedById;references:ID"`
}
