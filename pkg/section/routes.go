package section

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database

	sectionGroup := r.Group("/sections")
	{
		sectionGroup.POST("/create")
	}
}
