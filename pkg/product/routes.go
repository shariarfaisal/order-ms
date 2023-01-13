package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database

	group := r.Group("/product")
	{
		group.POST("/create", createProduct)
		group.GET("/", getProducts)
		// group.GET("/products/:id", getProductById)
		// group.PUT("/products/:id", updateProduct)
		// group.PATCH("/products/:id", updateProduct)
		// group.DELETE("/products/:id", deleteProduct)
	}

}