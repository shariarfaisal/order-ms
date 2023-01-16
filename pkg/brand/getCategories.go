package brand

import "github.com/gin-gonic/gin"

func getCategories(c *gin.Context) {
	categories, err := GetCategoryMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
