package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/category"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type ErrType map[string]interface{}

func getByIds(ids []int) []Brand {
	var brands []Brand
	db.Where("id IN ?", ids).Find(&brands)
	return brands
}

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

func getPartners(c *gin.Context) {
	partners, err := GetPartnerMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": partners})
}

func getBrands(c *gin.Context) {
	brands, err := GetBrandMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": brands})
}

type CreateCategorySchema struct {
	CategoryId uint   `json:"categoryId" v:"required"`
	BrandId    uint   `json:"brandId" v:"required"`
	Name       string `json:"name" v:"min=3;max=50"`
}

func createCategory(c *gin.Context) {
	var params CreateCategorySchema
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if isValid, err := validator.Validate(params); !isValid {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// is product category exists
	cat, err := category.GetCategoryById(params.CategoryId)
	if err != nil {
		c.JSON(400, gin.H{"error": ErrType{
			"categoryId": "Category not found",
		}})
		return
	}

	// is category already exists
	_, err = GetCategoryByCategoryId(params.CategoryId)
	if err == nil {
		c.JSON(400, gin.H{"error": ErrType{
			"categoryId": "Category already exists",
		}})
		return
	}

	// is brand exists
	brand, err := GetBrandById(params.BrandId)
	if err != nil {
		c.JSON(400, gin.H{"error": ErrType{
			"brandId": "Brand not found",
		}})
		return
	}

	categoryData := BrandCategory{
		CategoryId: params.CategoryId,
		BrandId:    params.BrandId,
	}

	if params.Name != "" {
		categoryData.Name = params.Name
	} else {
		categoryData.Name = cat.Name
	}

	categoryData.Slug = utils.GetSlug(categoryData.Name + "-" + brand.Name)

	err = db.Create(&categoryData).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": categoryData})
}

func getCategories(c *gin.Context) {
	categories, err := GetCategoryMany()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
