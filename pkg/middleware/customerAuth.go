package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
)

func CustomerAuth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is missing",
		})
		c.Abort()
		return
	}

	isValid, err, payloads := utils.ValidateJWT(strings.ReplaceAll(token, "Bearer ", ""), os.Getenv("APP_SECRET"))
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("customer", payloads)
	c.Next()
}
