package product

import "github.com/gin-gonic/gin"

func getProducts(c *gin.Context) {
	var products []Product
	if er := db.Find(&products).Error; er != nil {
		c.JSON(500, gin.H{"error": er.Error()})
		return
	}
	c.JSON(200, products)
}