package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/order-ms/pkg/validator"
)

type CreateAdminSchema struct {
	Name     string `json:"name" v:"required;min=3;max=50"`
	Email    string `json:"email" v:"required;email"`
	Phone    string `json:"phone" v:"required;min=11;max=11;"`
	Password string `json:"password" v:"required;min=6;"`
	Image    string `json:"image"`
	Role     string `json:"role" v:"required;"`
}

func createAdminUser(c *gin.Context) {
	_, exists := c.Get("admin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var payloads CreateAdminSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	adminUser := Admin{
		Name:     payloads.Name,
		Email:    payloads.Email,
		Phone:    payloads.Phone,
		Password: payloads.Password,
		Image:    payloads.Image,
		RoleName: payloads.Role,
	}

	_, err := GetRoleByName(payloads.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.ErrType{
				"role": "Role is not valid",
			},
		})
		return
	}

	emailExist, err := IsAdminExist("email", payloads.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if emailExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.ErrType{
				"email": "Email already exist",
			},
		})
		return
	}

	phoneExists, err := IsAdminExist("phone", payloads.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if phoneExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.ErrType{
				"phone": "Phone already exist",
			},
		})
		return
	}

	hashedPassword, err := utils.HashPassword(adminUser.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	adminUser.Password = hashedPassword

	err = db.Create(&adminUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin created successfully",
	})
}
