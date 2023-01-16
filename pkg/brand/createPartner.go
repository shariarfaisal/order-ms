package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type CreatePartnerSchema struct {
	Name     string `json:"name" v:"required;min=3;max=50"`
	UserName string `json:"userName" title:"User name" v:"required;min=3;max=50"`
	Email    string `json:"email" v:"required;email"`
	Phone    string `json:"phone" v:"required;phone"`
	Password string `json:"password" v:"required;min=6;max=50"`
}

func createPartner(c *gin.Context) {
	var params CreatePartnerSchema
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isValid, err := validator.Validate(params)
	if !isValid {
		c.JSON(400, gin.H{"error": err})
		return
	}

	partner := Partner{
		Name: params.Name,
	}

	// Create partner and partner user in a transaction
	er := db.Transaction(func(tx *gorm.DB) error {

		er := db.Create(&partner).Error
		if er != nil {
			return er
		}

		password := params.Password
		hashedPass, er := utils.HashPassword(password)
		if er != nil {
			return er
		}

		partnerUser := PartnerUser{
			PartnerId: partner.ID,
			Name:      params.UserName,
			Email:     params.Email,
			Phone:     params.Phone,
			Password:  hashedPass,
			Role:      PartnerRoleAdmin,
			Status:    PartnerUserActive,
		}

		er = db.Create(&partnerUser).Error
		if er != nil {
			return er
		}

		return nil
	})

	if er != nil {
		c.JSON(400, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": partner})
}
