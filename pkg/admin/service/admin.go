package service

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/admin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/order-ms/pkg/validator"
	"gorm.io/gorm"
)

type AdminService struct {
	Admin *admin.AdminRepo
	Role  *admin.RoleRepo
}

func NewUserService(db *gorm.DB) *AdminService {
	return &AdminService{
		Admin: admin.NewAdminRepo(db),
		Role:  admin.NewRoleRepo(db),
	}
}

type LoginAdminSchema struct {
	Email    string `json:"email" v:"required;email"`
	Password string `json:"password" v:"required;"`
}

func (s *AdminService) loginAdmin(c *gin.Context) {
	var payloads LoginAdminSchema
	c.ShouldBindJSON(&payloads)

	if isValid, errors := validator.Validate(payloads); !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	admin, err := s.Admin.GetByEmail(payloads.Email)
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

func (s *AdminService) getProfile(c *gin.Context) {
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

	admin, er := s.Admin.GetBy("id", id)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, admin)
}

type CreateAdminSchema struct {
	Name     string `json:"name" v:"required;min=3;max=50"`
	Email    string `json:"email" v:"required;email"`
	Phone    string `json:"phone" v:"required;min=11;max=11;"`
	Password string `json:"password" v:"required;min=6;"`
	Image    string `json:"image"`
	Role     string `json:"role" v:"required;"`
}

func (s *AdminService) createAdmin(c *gin.Context) {
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

	adminUser := admin.Admin{
		Name:     payloads.Name,
		Email:    payloads.Email,
		Phone:    payloads.Phone,
		Password: payloads.Password,
		Image:    payloads.Image,
		RoleName: payloads.Role,
	}

	_, err := s.Role.GetByName(payloads.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.ErrType{
				"role": "Role is not valid",
			},
		})
		return
	}

	emailExist, err := s.Admin.IsExist("email", payloads.Email)
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

	phoneExists, err := s.Admin.IsExist("phone", payloads.Phone)
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
