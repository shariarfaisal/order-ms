package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/middleware"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database

	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/create", middleware.AdminAuth, createAdminUser)
		adminGroup.POST("/login", loginAdminUser)
		adminGroup.GET("/me", middleware.AdminAuth, getProfile)
	}

	roleGroup := r.Group("/role")
	{
		roleGroup.POST("/create", createRole)
	}

}
