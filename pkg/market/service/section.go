package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/market"
	"gorm.io/gorm"
)

type SectionService struct {
	Store *market.StoreRepo
}

func NewSectionService(db *gorm.DB) *SectionService {
	return &SectionService{
		Store: market.NewStoreRepo(db),
	}
}

func (s *SectionService) create(c *gin.Context) {

}

func (s *SectionService) getItems(c *gin.Context) {

}

func (s *SectionService) getById(c *gin.Context) {

}

func (s *SectionService) update(c *gin.Context) {

}

func (s *SectionService) delete(c *gin.Context) {

}

func (s *SectionService) addItem(c *gin.Context) {

}

func (s *SectionService) updateItem(c *gin.Context) {

}

func (s *SectionService) deleteItem(c *gin.Context) {

}
