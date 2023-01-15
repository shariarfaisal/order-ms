package customer

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
)

type LoginSchema struct {
	Phone    string `json:"phone"`
	Password string `json:"password" v:"required;min=6;"`
	Email    string `json:"email"`
}

func login(c *gin.Context) {
	var payloads LoginSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	if payloads.Phone == "" && payloads.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone or Email is required",
		})
		return
	}

	var customer Customer
	if payloads.Phone != "" {
		data, err := GetBy("phone", payloads.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid phone or password",
			})
			return
		}
		customer = data
	} else {
		data, err := GetBy("email", payloads.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		customer = data
	}

	if !utils.IsValidPassword(payloads.Password, customer.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone or password",
		})
		return
	}

	jwtPayload := map[string]interface{}{
		"id":    customer.ID,
		"name":  customer.Name,
		"phone": customer.Phone,
		"email": customer.Email,
	}

	token, err := utils.GenerateJWT("customer", jwtPayload, os.Getenv("APP_SECRET"), 60*5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
