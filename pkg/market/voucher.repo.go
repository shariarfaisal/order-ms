package market

import "gorm.io/gorm"

type VoucherRepo struct {
	DB *gorm.DB
}

func NewVoucherRepo(db *gorm.DB) *VoucherRepo {
	return &VoucherRepo{DB: db}
}
