package brand

import "github.com/gin-gonic/gin"

func getPartners(c *gin.Context) {
	partners, err := GetPartnerMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": partners})
}
