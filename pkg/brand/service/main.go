package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	brand.Migration(db)

	partnerServices := NewPartnerService(db) // Partner Services

	partnerRoutes := r.Group("/partner")
	{
		partnerRoutes.POST("/create", partnerServices.createPartner)
		partnerRoutes.GET("/", partnerServices.getPartners)
		// r.GET("/:id", partnerServices.getPartner)
		// r.PUT("/:id", partnerServices.updatePartner)
		// r.DELETE("/:id", partnerServices.deletePartner)
	}

	bs := NewBrandService(db) // Brand Services
	brandRoutes := r.Group("/brand")
	{
		brandRoutes.POST("/create", bs.createBrand)
		brandRoutes.GET("/", bs.getBrands)

		bcs := NewBrandCategoryService(db) // Brand Category Services
		categoryRoutes := brandRoutes.Group("/category")
		{
			categoryRoutes.POST("/create", bcs.createBrandCategory)
			categoryRoutes.GET("/", bcs.getBrandCategories)
		}
	}

	productServices := NewProductService(db) // Product Services
	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/create", productServices.createProduct)
		productRoutes.GET("/", productServices.getProducts)
		// productRoutes.GET("/:id", productServices.getProductById)
		// productRoutes.PUT("/:id", productServices.updateProduct)
		// productRoutes.PATCH("/:id", productServices.updateProduct)
		productRoutes.DELETE("/:id", productServices.deleteProduct)
	}

	pcs := NewProductCategoryService(db) // Product Category Services
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", pcs.getProductCategories)
		categoryRoutes.POST("/create", pcs.createProductCategory)
		// categoryRoutes.GET("/categories/:id", pcs.getCategoryById)
		// categoryRoutes.PUT("/categories/:id", pcs.updateCategory)
		// categoryRoutes.DELETE("/categories/:id", pcs.deleteCategory)
	}

}
