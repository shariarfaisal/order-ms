package service

import (
	"github.com/shariarfaisal/order-ms/pkg/rider"
	"gorm.io/gorm"
)

type RiderService struct {
	Rider *rider.RiderRepo
}

func NewRiderService(db *gorm.DB) *RiderService {
	return &RiderService{
		Rider: rider.NewRiderRepo(db),
	}
}
