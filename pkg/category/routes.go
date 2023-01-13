package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database

	group := r.Group("/category")
	{
		group.GET("/", getCategories)
		group.POST("/create", createCategory)
		// group.GET("/categories/:id", getCategoryById)
		// group.PUT("/categories/:id", updateCategory)
		// group.DELETE("/categories/:id", deleteCategory)
	}

}