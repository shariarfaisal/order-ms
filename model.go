package main

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPending  OrderStatus = "pending"
	OrderAccepted OrderStatus = "accepted"
	OrderPreparing OrderStatus = "preparing"
	OrderReady OrderStatus = "ready"
	OrderDispatched OrderStatus = "dispatched"
	OrderDelivered OrderStatus = "delivered"
	OrderCancelled OrderStatus = "cancelled"
)

type Platform string

const (
	Web Platform = "web"
	Android Platform = "android"
	IOS Platform = "ios"
)

type OrderCharges struct {
	gorm.Model
	Total          int `json:"total" gorm:"<-create"`
	Discount       int `json:"discount" gorm:"<-create"`
	ServiceCharge  int `json:"service_charge" gorm:"<-create"`
	DeliveryCharge int `json:"delivery_charge" gorm:"<-create"`
	ItemsValue     int `json:"items_value"`
	OrderId int 	`json:"order_id" gorm:"<-create"`
	ItemDiscount int 	`json:"item_discount" gorm:"<-create"`
	PromoDiscount int 	`json:"promo_discount" gorm:"<-create"`
}

/*
	"id" INTEGER NOT NULL,
    "delivered_to" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "platform" VARCHAR(255) NULL,
    "dispatch_time" VARCHAR(255) NULL,
    "rider_note" VARCHAR(255) NULL,
    "confirmed_at" DATE NULL,
    "assigned_rider" INTEGER NULL,
    "hub_id" INTEGER NULL,
    "charges" INTEGER NULL,
    "edt" INTEGER NULL,
    "completed_at" DATE NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL,
    "payment_method" VARCHAR(255) NOT NULL,
    "payment_status" VARCHAR(255) NOT NULL
*/

type Order struct {
	gorm.Model
	DeliveredTo int `json:"delivered_to" gorm:"<-:create" `
	Status      OrderStatus `json:"status" gorm:"<-:create"`
	Platform    Platform `json:"platform" gorm:"<-:create"`
	DispatchTime string `json:"dispatched_time" gorm:"<-:create"`
	RiderNote string `json:"rider_note" gorm:"<-:create"`
	ConfirmedAt  time.Time	`json:"confirmed_at" gorm:"<-:create" gorm:"type:timestamp"`
	AssignedRider int `json:"assigned_rider" gorm:"<-:create" gorm:"index"`
	HubId   int `json:"hub_id" gorm:"<-:create" gorm:"index"`
	Charges int `json:"charges" gorm:"<-create"`
	EDT      int `json:"edt" gorm:"<-:create"`
	CompletedAt  time.Time `json:"completed_at" gorm:"<-:create" gorm:"type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"<-:create" gorm:"type:timestamp"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"<-:create" gorm:"type:timestamp"`
	PaymentMethod      string `json:"payment_method" gorm:"<-:create"`
	PaymentStatus   string `json:"payment_status" gorm:"<-:create"`
}

type OrderItem struct {
	gorm.Model
	ProductId int `json:"product_id" gorm:"<-:create" gorm:"index"`
	Quantity int `json:"quantity" gorm:"<-:create"`
	SaleUnit int `json:"sale_unit" gorm:"<-:create"`
	Total int `json:"total" gorm:"<-:create"`
	Discount int `json:"discount" gorm:"<-:create"`
	PickupId int `json:"pickup_id" gorm:"<-:create" gorm:"index"`
	OrderId int `json:"order_id" gorm:"<-:create" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:create" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"<-:create" gorm:"type:timestamp"`
}

type PickupStatus string

const (
	PickupAccepted PickupStatus = "accepted"
	PickupRejected PickupStatus = "rejected"
)

type Pickup struct {
	gorm.Model 
	BrandId int `json:"brand_id" gorm:"<-:create" gorm:"index"`
	OrderId int `json:"order_id" gorm:"<-:create" gorm:"index"`
	Note string `json:"note" gorm:"<-:create"`
	Total int `json:"total" gorm:"<-:create"`
	Status PickupStatus `json:"status" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:create" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"<-:create" gorm:"type:timestamp"`
}


type DeliveryAddress struct {
	gorm.Model
	Name string `json:"name" gorm:"<-:create"`
	Phone string `json:"phone" gorm:"<-:create" gorm:"index"`
	Address string `json:"address" gorm:"<-:create"`
	Area string `json:"area" gorm:"<-:create"`
	Lat float64 `json:"lat" gorm:"<-:create"`
	Lng float64 `json:"lng" gorm:"<-:create"`
	OrderId int `json:"order_id" gorm:"<-:create" gorm:"index"`
	UserId int `json:"user_id" gorm:"<-:create" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:create" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"<-:create" gorm:"type:timestamp"`
}

type Geo struct {
	Lat float64
	Lng float64
}

type AssignedRider struct {
	gorm.Model
	OrderId int
	RiderId int
	Movements []Geo
}

type OrderPaymentMethod string
const ( 
	PaymentMethodCash OrderPaymentMethod = "cash"
	PaymentMethodCard OrderPaymentMethod = "card"
	PaymentMethodWallet OrderPaymentMethod = "wallet"
)

type OrderPaymentStatus string
const (
	PaymentStatusPending OrderPaymentStatus = "pending"
	PaymentStatusSuccess OrderPaymentStatus = "success"
	PaymentStatusFailed OrderPaymentStatus = "failed"
)

type OrderPayment struct {
	gorm.Model
	Method OrderPaymentMethod
	Status OrderPaymentStatus
	OrderId int
}

type PaymentLog struct {
	gorm.Model
	OrderId int
	Method OrderPaymentMethod
	TrxId string
	Amount float64
	Status OrderPaymentStatus
}

type Voucher struct {
	gorm.Model
	Amount float64
	Type string
	Code string
}

type RiderLocation struct {
	Lat float64
	Lng float64
	RiderId int
}

type Rider struct {
	gorm.Model
	Name string
	Phone string
	Image string
	Type string
	Status string
}

type Customer struct {
	gorm.Model
	Name string
	Image string
	Phone string
	Username string
	Email string
}

type Timeline struct {
	gorm.Model
	OrderId int
	Message string
	ActionType string
	ReferenceId int
}

type CartItem struct {
	gorm.Model
	UserId int
	ProductId int
	Quantity int
}