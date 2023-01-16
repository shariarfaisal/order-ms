package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Brand "github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
)

type CreateCategorySchema struct {
	CategoryId uint   `json:"categoryId" v:"required"`
	BrandId    uint   `json:"brandId" v:"required"`
	Name       string `json:"name" v:"min=3;max=50"`
}

func createBrandCategory(c *gin.Context) {
	pCategoryRepo := Brand.NewProductCategoryRepo(db)
	bCategoryRepo := Brand.NewBrandCategoryRepo(db)
	brandRepo := Brand.NewBrandRepo(db)

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
	cat, err := pCategoryRepo.GetById(params.CategoryId)
	if err != nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category not found",
		}})
		return
	}

	// is category already exists
	_, err = bCategoryRepo.GetByPCategoryId(params.CategoryId)
	if err == nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category already exists",
		}})
		return
	}

	// is brand exists
	brand, err := brandRepo.GetById(params.BrandId)
	if err != nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"brandId": "Brand not found",
		}})
		return
	}

	categoryData := Brand.BrandCategory{
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

func getBrandCategories(c *gin.Context) {
	bCategoryRepo := Brand.NewBrandCategoryRepo(db)

	categories, err := bCategoryRepo.GetItems()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
