package hub

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database
	hr := r.Group("/hub")
	{
		hr.GET("/", getMany)
		hr.GET("/:id", getById)
		hr.POST("/create", createHub)
		// hr.PUT("/:id", updateHub)
		// hr.DELETE("/:id", deleteHub)
	}
}