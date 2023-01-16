package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProfile(c *gin.Context) {
	payload, err := c.Get("admin")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	id, isId := payload.(map[string]interface{})["id"]
	if !isId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	admin, er := GetBy("id", id)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, admin)
}
