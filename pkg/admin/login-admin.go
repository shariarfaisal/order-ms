package admin

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/order-ms/pkg/validator"
)

type LoginAdminSchema struct {
	Email    string `json:"email" v:"required;email"`
	Password string `json:"password" v:"required;"`
}

func loginAdminUser(c *gin.Context) {
	var payloads LoginAdminSchema
	c.ShouldBindJSON(&payloads)

	if isValid, errors := validator.Validate(payloads); !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	admin, err := GetAdminByEmail(payloads.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email. or password",
		})
		return
	}

	if !utils.IsValidPassword(payloads.Password, admin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password.",
		})
		return
	}

	tokenPayload := map[string]interface{}{
		"id":    admin.ID,
		"name":  admin.Name,
		"email": admin.Email,
		"phone": admin.Phone,
		"role":  admin.RoleName,
		"image": admin.Image,
	}
	token, err := utils.GenerateJWT("bearer", tokenPayload, os.Getenv("APP_SECRET"), 60*2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
