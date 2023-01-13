package customer

import (
	"time"

	"gorm.io/gorm"
)

type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
	CustomerStatusBlocked  CustomerStatus = "blocked"
)

type CustomerGender string
const ( 
	CustomerGenderMale CustomerGender = "male" 
	CustomerGenderFemale CustomerGender = "female"
)

type Customer struct {
	ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
	Name     string `json:"name" gorm:"<-:create"`
	Image    string `json:"image" gorm:"<-:create"`
	Phone    string `json:"phone" gorm:"<-:create" gorm:"index"`
	Username string `json:"username" gorm:"<-:create" gorm:"index"`
	Email    string `json:"email" gorm:"<-:create" gorm:"index"`
	EmailVerified bool `json:"email_verified" gorm:"<-:create"`
	PushToken string `json:"push_token" gorm:"<-:create"`
	Password string `json:"password" gorm:"<-:create"`
	IsActive bool `json:"is_active" gorm:"<-:create"`
	Status CustomerStatus `json:"status" gorm:"<-:create"`
	LastLogin time.Time `json:"last_login" gorm:"<-:create;type:timestamp"`
	Gender CustomerGender `json:"gender" gorm:"<-:create"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"<-:create;type:date"`
}

type CustomerAddress struct {
	ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
	CustomerId int `json:"customerId" gorm:"<-:create;index"`
	Customer Customer `json:"customer" gorm:"<-:create;foreignKey:CustomerId;references:ID"`
	Label string `json:"label" gorm:"<-:create"`
	Address string `json:"address" gorm:"<-:create"`
	Area string `json:"area" gorm:"<-:create"`
	PostalCode string `json:"postal_code" gorm:"<-:create"`
	Latitude float64 `json:"latitude" gorm:"<-:create"`
	Longitude float64 `json:"longitude" gorm:"<-:create"`
	Apartment string `json:"apartment" gorm:"<-:create"`
	Floor string `json:"floor" gorm:"<-:create"`
	RoadNo string `json:"road_no" gorm:"<-:create"`
	
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&CustomerAddress{})
}