package hub

import "gorm.io/gorm"

type Hub struct {
	gorm.Model
	Name        string `json:"name" gorm:"<-:create"`
	Slug        string `json:"slug" gorm:"<-:create"`
	City 	  string `json:"city" gorm:"<-:create"`
	Area 	  string `json:"area" gorm:"<-:create"`
	Region 	  string `json:"region" gorm:"<-:create"`
	Country 	  string `json:"country" gorm:"<-:create"`
	Latitude 	float64 `json:"latitude" gorm:"<-:create"`
	Longitude 	float64 `json:"longitude" gorm:"<-:create"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Hub{})
}