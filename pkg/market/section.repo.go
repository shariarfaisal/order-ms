package market

import "gorm.io/gorm"

type SectionRepo struct {
	DB *gorm.DB
}

func NewSectionRepo(db *gorm.DB) *SectionRepo {
	return &SectionRepo{DB: db}
}
