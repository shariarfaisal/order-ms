package hub

import "gorm.io/gorm"

type HubRepo struct {
	DB *gorm.DB
}

func NewHubRepo(db *gorm.DB) *HubRepo {
	return &HubRepo{DB: db}
}

func (r *HubRepo) GetById(id uint) (Hub, error) {
	var hub Hub
	result := r.DB.First(&hub, id)
	return hub, result.Error
}

func (r *HubRepo) GetItems() ([]Hub, error) {
	var hubs []Hub
	result := r.DB.Find(&hubs)
	return hubs, result.Error
}
