package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
)

func AdminAuth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
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
	fmt.Println("valid", payloads)

	c.Set("admin", payloads)
	c.Next()
}
