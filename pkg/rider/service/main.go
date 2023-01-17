package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/rider"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&rider.Rider{})
}

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)
}
