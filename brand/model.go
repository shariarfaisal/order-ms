package brand

import (
	"github.com/shariarfaisal/order-ms/hub"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name        string `json:"name" gorm:"<-:create"`
	Slug        string `json:"slug" gorm:"<-:create"`
	Details     string `json:"details" gorm:"<-:create"`
	Phone       string `json:"phone" gorm:"<-:create"`
	Email       string `json:"email" gorm:"<-:create"`
	Status 	string `json:"status" gorm:"<-:create"`
	OperatingDays []int `json:"operating_days" gorm:"<-:create"`
	OperatingHours []int `json:"operating_hours" gorm:"<-:create"`
	HubId int `json:"hub_id" gorm:"<-:create;index"`
	Hub hub.Hub `json:"hub" gorm:"<-:create;foreignKey:HubId"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Brand{})
}