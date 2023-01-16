package service

import (
	"github.com/shariarfaisal/order-ms/pkg/market"
	"gorm.io/gorm"
)

type StoreService struct {
	Store *market.StoreRepo
}

func NewStoreService(db *gorm.DB) *StoreService {
	return &StoreService{
		Store: market.NewStoreRepo(db),
	}
}
