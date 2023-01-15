package section

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	sectionGroup := r.Group("/sections")
	{
		sectionGroup.POST("/create")
	}
}
