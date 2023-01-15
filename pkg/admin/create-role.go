package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/validator"
)

type CreateRoleSchema struct {
	Name string `json:"name" title:"Name" v:"required;min=3;max=50"`
}

func createRole(c *gin.Context) {
	var payloads CreateRoleSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	role := Role{
		Name: payloads.Name,
	}

	err := db.Create(&role).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role created successfully",
	})
}
