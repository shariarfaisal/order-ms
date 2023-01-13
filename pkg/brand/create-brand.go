package brand

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type CreateBrandSchema struct {
	Name        string    `json:"name" v:"required;min=3;max=50"`
	Type        BrandType `json:"type" v:"required;enum=store,restaurant,grocery"`
	Details     string    `json:"details"`
	Phone       string    `json:"phone" v:"required;phone"`
	Email       string    `json:"email" v:"required;email"`
	Logo        string    `json:"logo" v:"required"`
	BannerImage string    `json:"bannerImage" v:"required"`
	PartnerId   uint      `json:"partnerId" v:"required"`
	Address     struct {
		Address    string  `json:"address" v:"required"`
		Area       string  `json:"area" v:"required"`
		PostalCode string  `json:"postalCode" title:"Postal code" v:"required;eq='4'"`
		Latitude   float64 `json:"latitude" v:"required"`
		Longitude  float64 `json:"longitude" v:"required"`
		Apartment  string  `json:"apartment"`
		Floor      string  `json:"floor"`
		RoadNo     string  `json:"roadNo" title:"Road no"`
	} `json:"address" v:"required"`
	OperatingTimes map[string][]struct {
		From struct {
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"from" v:"required"`
		To struct {
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"to" v:"required"`
	} `json:"operatingTimes"`
	HubId uint `json:"hubId" v:"required"`
}

func createBrand(c *gin.Context) {
	var params CreateBrandSchema

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := validator.Validate(params)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, hubErr := hub.GetById(params.HubId)
	if hubErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]string{"HubId": "Hub not found"}})
		return
	}

	_, partnerExists := GetPartnerById(params.PartnerId)
	if partnerExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]string{"PartnerId": "Partner not found"}})
		return
	}

	tx := db.Select(" name ").Where("hub_id = ? AND name = ?", params.HubId, params.Name).First(&Brand{})
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]string{"Name": "Brand name already exists"}})
		return
	}

	// copy params address to brand address
	address := BrandAddress{
		Address:    params.Address.Address,
		Area:       params.Address.Area,
		PostalCode: params.Address.PostalCode,
		Latitude:   params.Address.Latitude,
		Longitude:  params.Address.Longitude,
		Apartment:  params.Address.Apartment,
		Floor:      params.Address.Floor,
		RoadNo:     params.Address.RoadNo,
	}

	brand := Brand{
		Name:          params.Name,
		Slug:          utils.GetSlug(params.Name),
		Type:          params.Type,
		Details:       params.Details,
		Phone:         params.Phone,
		Email:         params.Email,
		EmailVerified: false,
		Logo:          params.Logo,
		BannerImage:   params.BannerImage,
		Rating:        5,
		PartnerId:     params.PartnerId,
		Status:        BrandStatusPending,
		IsAvailable:   false,
		AddressId:     0,
		HubId:         params.HubId,
	}

	if params.OperatingTimes != nil {
		operatingTime := map[string]interface{}{}
		for day, times := range params.OperatingTimes {
			operatingTime[day] = []OperatingTime{}
			for _, time := range times {
				operatingTime[day] = append(operatingTime[day].([]OperatingTime), OperatingTime{
					From: time.From,
					To:   time.To,
				})
			}
		}

		brand.OperatingTimes = operatingTime
	}

	er := db.Transaction(func(tx *gorm.DB) error {
		// create brand address
		if err := tx.Create(&address).Error; err != nil {
			return err
		}

		// set brand address id
		brand.AddressId = address.ID

		// create brand
		if err := tx.Create(&brand).Error; err != nil {
			return err
		}

		return nil
	})

	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": Brand{
		ID:          brand.ID,
		Name:        brand.Name,
		Slug:        brand.Slug,
		Type:        brand.Type,
		Details:     brand.Details,
		Logo:        brand.Logo,
		BannerImage: brand.BannerImage,
		Address:     address,
	}})
}
