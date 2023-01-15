package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

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

func getProducts(c *gin.Context) {
	var products []Product
	if er := db.Find(&products).Error; er != nil {
		c.JSON(500, gin.H{"error": er.Error()})
		return
	}
	c.JSON(200, products)
}
