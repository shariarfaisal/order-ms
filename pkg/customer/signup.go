package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
)

type SignUpSchema struct {
	Name        string `json:"name" v:"required;min=3;"`
	Email       string `json:"email" v:"email"`
	Phone       string `json:"phone" v:"required;min=11;max=11;"`
	Password    string `json:"password" v:"required;min=6;"`
	Image       string `json:"image"`
	Gender      string `json:"gender" v:"enum=male|female"`
	DateOfBirth string `json:"dateOfBirth" v:"date"`
}

func signUp(c *gin.Context) {
	var payloads SignUpSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	/*
		is email exist
		is phone exist
		hash password
	*/

	if exist, err := IsCustomerExist("phone", payloads.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone already exist",
		})
		return
	}

	if exist, err := IsCustomerExist("email", payloads.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exist",
		})
		return
	}

	customer := Customer{
		Name:     payloads.Name,
		Email:    payloads.Email,
		Phone:    payloads.Phone,
		Password: payloads.Password,
		Image:    payloads.Image,
		Gender:   CustomerGender(payloads.Gender),
		Status:   CustomerStatusActive,
		IsActive: true,
	}

	if payloads.DateOfBirth != "" {
		dateOfBirth, err := utils.ParseDate(payloads.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		customer.DateOfBirth = dateOfBirth
	}

	hashPass, err := utils.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.Password = hashPass

	err = db.Create(&customer).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
	})
}
