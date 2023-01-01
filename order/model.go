package order

import (
	"time"

	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/customer"
	"github.com/shariarfaisal/order-ms/hub"
	"github.com/shariarfaisal/order-ms/product"
	"github.com/shariarfaisal/order-ms/rider"
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

type Platform string

const (
	Web     Platform = "web"
	Android Platform = "android"
	IOS     Platform = "ios"
)

type Order struct {
	gorm.Model
	DeliveredTo   int         `json:"deliveredTo" gorm:"<-:create"`
	DeliveryAddress DeliveryAddress `json:"deliveryAddress" gorm:"foreignKey:DeliveredTo;references:ID"`
	Status        OrderStatus `json:"status" gorm:"<-:create"`
	Platform      Platform    `json:"platform" gorm:"<-:create"`
	DispatchTime  string      `json:"dispatchedTime" gorm:"<-:create"`
	RiderNote     string      `json:"riderNote" gorm:"<-:create"`
	ConfirmedAt   time.Time   `json:"confirmedAt" gorm:"<-:create" gorm:"type:timestamp"`
	AssignedTo int         `json:"assignedTo" gorm:"<-:create" gorm:"index"`
	AssignedRider AssignedRider `json:"assignedRider" gorm:"foreignKey:AssignedTo;references:ID"`
	HubId         int         `json:"hubId" gorm:"<-:create" gorm:"index"`
	Hub 			hub.Hub `json:"hub" gorm:"foreignKey:HubId"`
	EDT           int         `json:"edt" gorm:"<-:create"`
	PaymentMethod string      `json:"paymentMethod" gorm:"<-:create"`
	PaymentStatus string      `json:"paymentStatus" gorm:"<-:create"`
	ChargesId    int         `json:"chargesId" gorm:"<-:create" gorm:"index"`
	Charges      OrderCharges `json:"charges" gorm:"foreignKey:ChargesId;references:ID"`
	Pickups []Pickup `json:"pickups" gorm:"foreignKey:OrderId;references:ID"`
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderId;references:ID"`
	CompletedAt   time.Time   `json:"completedAt" gorm:"<-:create" gorm:"type:timestamp"`
}

type OrderCharges struct {
	gorm.Model
	Total 	   int         		`json:"total" gorm:"<-create"`
	ProductDiscount      int         `json:"productDiscount" gorm:"<-create"`
	VoucherDiscount 	int         `json:"voucherDiscount" gorm:"<-create"`
	Voucher struct{
		Code string `json:"code" gorm:"<-create"`
		DiscountAmount int `json:"discountAmount" gorm:"<-create"`
		DiscountType string `json:"discountType" gorm:"<-create"`
	}         `json:"voucher" gorm:"<-create"`
	DeliveryCharge int         `json:"deliveryCharge" gorm:"<-create"`
	DeliveryDiscount int         `json:"deliveryDiscount" gorm:"<-create"`
	TotalDiscount int         `json:"totalDiscount" gorm:"<-create"`
	ServiceCharge int         `json:"serviceCharge" gorm:"<-create"`
	SubTotal int         `json:"subTotal" gorm:"<-create"`
}

type PickupStatus string

const (
	PickupAccepted PickupStatus = "accepted"
	PickupRejected PickupStatus = "rejected"
)

type Pickup struct {
	gorm.Model
	BrandId   int          `json:"brand_id" gorm:"<-:create" gorm:"index"`
	Brand    brand.Brand `json:"brand" gorm:"foreignKey:BrandId"`
	OrderId   int          `json:"order_id" gorm:"<-:create" gorm:"index"`
	Order    Order        `json:"order" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // OrderId is the foreign key
	OrderNumber string     `json:"order_number" gorm:"<-:create"`
	Note      string       `json:"note" gorm:"<-:create"`
	Total     int          `json:"total" gorm:"<-:create"`
	Status    PickupStatus `json:"status" gorm:"<-:create"`
	Items []OrderItem `json:"items" gorm:"foreignKey:PickupId;references:ID"`
}

type OrderItem struct {
	gorm.Model
	ProductId int       `json:"product_id" gorm:"<-:create" gorm:"index"`
	Quantity  int       `json:"quantity" gorm:"<-:create"`
	SaleUnit  int       `json:"sale_unit" gorm:"<-:create"`
	Total     int       `json:"total" gorm:"<-:create"`
	Discount  int       `json:"discount" gorm:"<-:create"`
	PickupId  int       `json:"pickup_id" gorm:"<-:create" gorm:"index"`
	Pickup    Pickup    `json:"pickup" gorm:"foreignKey:PickupId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderId   int       `json:"order_id" gorm:"<-:create" gorm:"index"`
	Order    Order        `json:"order" gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product product.Product `json:"product" gorm:"<-:create;foreignKey:ProductId;references:ID"`
}

type DeliveryAddress struct {
	gorm.Model
	Name      string    `json:"name" gorm:"<-:create"`
	Phone     string    `json:"phone" gorm:"<-:create" gorm:"index"`
	Address   string    `json:"address" gorm:"<-:create"`
	Area      string    `json:"area" gorm:"<-:create"`
	Lat       float64   `json:"lat" gorm:"<-:create"`
	Lng       float64   `json:"lng" gorm:"<-:create"`
	CustomerId    int       `json:"customer_id" gorm:"<-:create" gorm:"index"`
	Customer 	customer.Customer `json:"customer" gorm:"<-:create;foreignKey:CustomerId;references:ID"`
}

type AssignedRider struct {
	gorm.Model
	RiderId   int
	Rider rider.Rider `json:"rider" gorm:"<-:create;foreignKey:RiderId;references:ID"`
	Note 	string `json:"note" gorm:"<-:create"`
}


func Migration(db *gorm.DB) {
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Pickup{})
	db.AutoMigrate(&OrderItem{})
	db.AutoMigrate(&DeliveryAddress{})
	db.AutoMigrate(&AssignedRider{})
}