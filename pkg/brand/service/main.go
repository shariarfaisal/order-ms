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

	partnerRoutes := r.Group("/partner")
	{
		partnerRoutes.POST("/create", createPartner)
		partnerRoutes.GET("/", getPartners)
		// r.GET("/:id", getPartner)
		// r.PUT("/:id", updatePartner)
		// r.DELETE("/:id", deletePartner)
	}

	brandRoutes := r.Group("/brand")
	{
		brandRoutes.POST("/create", createBrand)
		brandRoutes.GET("/", getBrands)
		categoryRoutes := brandRoutes.Group("/category")
		{
			categoryRoutes.POST("/create", createBrandCategory)
			categoryRoutes.GET("/", getBrandCategories)
		}
	}

	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/create", createProduct)
		productRoutes.GET("/", getProducts)
		// productRoutes.GET("/:id", getProductById)
		// productRoutes.PUT("/:id", updateProduct)
		// productRoutes.PATCH("/:id", updateProduct)
		productRoutes.DELETE("/:id", deleteProduct)
	}

	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.GET("/", getProductCategories)
		categoryRoutes.POST("/create", createProductCategory)
		// categoryRoutes.GET("/categories/:id", getCategoryById)
		// categoryRoutes.PUT("/categories/:id", updateCategory)
		// categoryRoutes.DELETE("/categories/:id", deleteCategory)
	}

}
