package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/category"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
)

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
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category not found",
		}})
		return
	}

	// is category already exists
	_, err = GetCategoryByCategoryId(params.CategoryId)
	if err == nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category already exists",
		}})
		return
	}

	// is brand exists
	brand, err := GetBrandById(params.BrandId)
	if err != nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
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
