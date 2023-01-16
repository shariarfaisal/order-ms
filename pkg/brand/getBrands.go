package brand

import "github.com/gin-gonic/gin"

func getBrands(c *gin.Context) {
	brands, err := GetBrandMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": brands})
}
