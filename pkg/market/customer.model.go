package market

import (
	"time"
)

type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
	CustomerStatusBlocked  CustomerStatus = "blocked"
)

type CustomerGender string

const (
	CustomerGenderMale   CustomerGender = "male"
	CustomerGenderFemale CustomerGender = "female"
)

type CustomerJWTPayload struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Image string `json:"image"`
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
}

type Customer struct {
	ID            uint           `gorm:"primarykey"`
	CreatedAt     time.Time      `json:"createdAt,omitempty" gorm:"<-:create;type:timestamp"`
	UpdatedAt     time.Time      `json:"updatedAt,omitempty" gorm:"<-:create;type:timestamp"`
	Name          string         `json:"name,omitempty" gorm:"<-:create"`
	Image         string         `json:"image,omitempty" gorm:"<-:create"`
	Phone         string         `json:"phone,omitempty" gorm:"<-:create" gorm:"index"`
	Email         string         `json:"email,omitempty" gorm:"<-:create" gorm:"index"`
	EmailVerified bool           `json:"emailVerified,omitempty" gorm:"<-:create"`
	PushToken     string         `json:"pushToken,omitempty" gorm:"<-:create"`
	Password      string         `json:"-" gorm:"<-:create"`
	IsActive      bool           `json:"isActive,omitempty" gorm:"<-:create"`
	Status        CustomerStatus `json:"status,omitempty" gorm:"<-:create"`
	LastLogin     time.Time      `json:"lastLogin,omitempty" gorm:"<-:create;type:timestamp"`
	Gender        CustomerGender `json:"gender,omitempty" gorm:"<-:create"`
	DateOfBirth   time.Time      `json:"dateOfBirth,omitempty" gorm:"<-:create;type:date"`
}

type CustomerAddress struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerId int      `json:"customerId" gorm:"<-:create;index"`
	Customer   Customer `json:"customer" gorm:"<-:create;foreignKey:CustomerId;references:ID"`
	Label      string   `json:"label" gorm:"<-:create"`
	Address    string   `json:"address" gorm:"<-:create"`
	Area       string   `json:"area" gorm:"<-:create"`
	PostalCode string   `json:"postal_code" gorm:"<-:create"`
	Latitude   float64  `json:"latitude" gorm:"<-:create"`
	Longitude  float64  `json:"longitude" gorm:"<-:create"`
	Apartment  string   `json:"apartment" gorm:"<-:create"`
	Floor      string   `json:"floor" gorm:"<-:create"`
	RoadNo     string   `json:"road_no" gorm:"<-:create"`
}
