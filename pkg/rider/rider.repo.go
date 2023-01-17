package rider

import "gorm.io/gorm"

type RiderRepo struct {
	DB *gorm.DB
} 

func NewRiderRepo(db *gorm.DB) *RiderRepo {
	return &RiderRepo{DB: db}
}
			
			