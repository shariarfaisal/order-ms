package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	customerGroup := r.Group("/customer")
	{
		customerGroup.POST("/signup", signUp)
		customerGroup.POST("/login", login)
		customerGroup.GET("/me", middleware.CustomerAuth, getProfile)
	}
}
