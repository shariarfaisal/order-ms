package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Brand "github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type BrandCategoryService struct {
	ProductCategory *Brand.ProductCategoryRepo
	BrandCategory   *Brand.BrandCategoryRepo
	Brand           *Brand.BrandRepo
}

func NewBrandCategoryService(db *gorm.DB) *BrandCategoryService {
	return &BrandCategoryService{
		ProductCategory: Brand.NewProductCategoryRepo(db),
		BrandCategory:   Brand.NewBrandCategoryRepo(db),
		Brand:           Brand.NewBrandRepo(db),
	}
}

type CreateCategorySchema struct {
	CategoryId uint   `json:"categoryId" v:"required"`
	BrandId    uint   `json:"brandId" v:"required"`
	Name       string `json:"name" v:"min=3;max=50"`
}

func (s *BrandCategoryService) createBrandCategory(c *gin.Context) {
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
	cat, err := s.ProductCategory.GetById(params.CategoryId)
	if err != nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category not found",
		}})
		return
	}

	// is category already exists
	_, err = s.BrandCategory.GetByPCategoryId(params.CategoryId)
	if err == nil {
		c.JSON(400, gin.H{"error": utils.ErrType{
			"categoryId": "Category already exists",
		}})
		return
	}

	// is brand exists
	brand, err := s.Brand.GetById(params.BrandId)
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

func (s *BrandCategoryService) getBrandCategories(c *gin.Context) {
	categories, err := s.BrandCategory.GetItems()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
