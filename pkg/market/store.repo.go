package market

import "gorm.io/gorm"

type StoreRepo struct {
	DB *gorm.DB
}

func NewStoreRepo(db *gorm.DB) *StoreRepo {
	return &StoreRepo{DB: db}
}
