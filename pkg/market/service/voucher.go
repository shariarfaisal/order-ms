package service

import (
	"github.com/shariarfaisal/order-ms/pkg/market"
	"gorm.io/gorm"
)

type VoucherService struct {
	Voucher *market.VoucherRepo
}

func NewVoucherService(db *gorm.DB) *VoucherService {
	return &VoucherService{
		Voucher: market.NewVoucherRepo(db),
	}
}
