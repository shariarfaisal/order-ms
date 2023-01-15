package category

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	group := r.Group("/category")
	{
		group.GET("/", getCategories)
		group.POST("/create", createCategory)
		// group.GET("/categories/:id", getCategoryById)
		// group.PUT("/categories/:id", updateCategory)
		// group.DELETE("/categories/:id", deleteCategory)
	}

}

type CategorySchema struct {
	Name  string `json:"name" v:"required;min=3;max=50"`
	Icon  string `json:"icon" v:"required"`
	Image string `json:"image" v:"required"`
}

func createCategory(c *gin.Context) {
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

	category := ProductCategory{
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

	c.JSON(200, gin.H{"result": category})
}

func getCategories(c *gin.Context) {
	categories, err := GetCategories()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": categories})
}
