package brand

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database

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
			categoryRoutes.POST("/create", createCategory)
			categoryRoutes.GET("/", getCategories)
		}
	}

}