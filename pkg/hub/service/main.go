package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migration(db *gorm.DB) {
	db.AutoMigrate(&hub.Hub{})
}

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	hubServices := NewHubService(db)
	hr := r.Group("/hubs")
	{
		hr.GET("/", hubServices.getMany)
		hr.GET("/:id", hubServices.getById)
		hr.POST("/create", hubServices.createHub)
		// hr.PUT("/:id", updateHub)
		// hr.DELETE("/:id", deleteHub)
	}
}
