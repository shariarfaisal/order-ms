package service

import (
	"github.com/gin-gonic/gin"
	Brand "github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/order-ms/pkg/validator"
)

type CategorySchema struct {
	Name  string `json:"name" v:"required;min=3;max=50"`
	Icon  string `json:"icon" v:"required"`
	Image string `json:"image" v:"required"`
}

func createProductCategory(c *gin.Context) {
	pCategoryRepo := Brand.NewProductCategoryRepo(db)

	var params CategorySchema
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isValid, err := validator.Validate(params)
	if !isValid {
		c.JSON(400, gin.H{"error": err})
		return
	}

	exists, isErr := pCategoryRepo.IsExists("name", params.Name)
	if isErr != nil {
		c.JSON(400, gin.H{"error": isErr.Error()})
		return
	}

	if exists {
		c.JSON(400, gin.H{"error": "Category already exists"})
		return
	}

	category := Brand.ProductCategory{
		Name:  params.Name,
		Slug:  utils.GetSlug(params.Name),
		Icon:  params.Icon,
		Image: params.Image,
	}

	er := db.Create(&category).Error
	if er != nil {
		c.JSON(400, gin.H{"error": er.Error()})
		return
	}

	c.JSON(200, category)
}

func getProductCategories(c *gin.Context) {
	pCategoryRepo := Brand.NewProductCategoryRepo(db)
	categories, err := pCategoryRepo.GetItems()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
