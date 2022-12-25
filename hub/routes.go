package hub

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database
	// hub := r.Group("/hub")
	// {
	// 	hub.GET("/", getHubs)
	// 	hub.GET("/:id", getHub)
	// 	hub.POST("/", createHub)
	// 	hub.PUT("/:id", updateHub)
	// 	hub.DELETE("/:id", deleteHub)
	// }
}