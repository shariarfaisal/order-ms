package brand

import (
	"encoding/json"
	"errors"

	"github.com/shariarfaisal/order-ms/pkg/hub"
	"gorm.io/gorm"
)

type BrandStatus string

const (
	BrandStatusPending  BrandStatus = "pending"
	BrandStatusActive   BrandStatus = "active"
	BrandStatusInactive BrandStatus = "inactive"
)

type BrandType string

const (
	BrandTypeStore      BrandType = "store"
	BrandTypeRestaurant BrandType = "restaurant"
	BrandTypeGrocery    BrandType = "grocery"
)

type OperatingTime struct {
	From struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"from,omitempty"`
	To struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
	} `json:"to,omitempty"`
}

type OperatingTimes map[string]interface{}

func (p *OperatingTimes) Scan(value interface{}) error {
	source, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(OperatingTimes)
	if !ok {
		return errors.New("Type assertion OperatingTimes failed.")
	}

	return nil
}

type Brand struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	Type           BrandType      `json:"type"`
	Details        string         `json:"details"`
	Phone          string         `json:"phone"`
	Email          string         `json:"email"`
	EmailVerified  bool           `json:"emailVerified"`
	Logo           string         `json:"logo"`
	BannerImage    string         `json:"bannerImage"`
	Rating         float32        `json:"rating"`
	PartnerId      uint           `json:"partnerId"`
	Partner        Partner        `json:"-" gorm:"foreignKey:PartnerId"`
	Prefix         string         `json:"prefix"`
	Status         BrandStatus    `json:"status"`
	IsAvailable    bool           `json:"isAvailable"`
	AddressId      uint           `json:"addressId" gorm:"index"`
	Address        BrandAddress   `json:"-" gorm:"foreignKey:AddressId"`
	OperatingTimes OperatingTimes `json:"operatingTimes,omitempty" gorm:"type:jsonb"`
	HubId          uint           `json:"hubId" gorm:"index"`
	Hub            hub.Hub        `json:"-" gorm:"foreignKey:HubId"`
}

type BrandAddress struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Address    string  `json:"address"`
	Area       string  `json:"area"`
	PostalCode string  `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Apartment  string  `json:"apartment"`
	Floor      string  `json:"floor"`
	RoadNo     string  `json:"road_no"`
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&Brand{})
	db.AutoMigrate(&BrandAddress{})
	db.AutoMigrate(&BrandCategory{})
	db.AutoMigrate(&Partner{})
	db.AutoMigrate(&PartnerUser{})
}
