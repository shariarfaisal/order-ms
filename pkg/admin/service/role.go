package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/admin"
	"github.com/shariarfaisal/order-ms/pkg/validator"
	"gorm.io/gorm"
)

type RoleService struct {
	Role *admin.RoleRepo
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		Role: admin.NewRoleRepo(db),
	}
}

type CreateRoleSchema struct {
	Name string `json:"name" title:"Name" v:"required;min=3;max=50"`
}

func (s *RoleService) createRole(c *gin.Context) {
	var payloads CreateRoleSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	role := admin.Role{
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
