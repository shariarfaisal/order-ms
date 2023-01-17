package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"github.com/shariarfaisal/order-ms/pkg/validator"
	"gorm.io/gorm"
)

type HubService struct {
	Hub *hub.HubRepo
}

func NewHubService(db *gorm.DB) *HubService {
	return &HubService{
		Hub: hub.NewHubRepo(db),
	}
}

type createDto struct {
	Name      string  `json:"name" v:"required"`
	City      string  `json:"city" v:"required"`
	Area      string  `json:"area" v:"required"`
	Region    string  `json:"region" v:"required"`
	Country   string  `json:"country" v:"required"`
	Latitude  float64 `json:"latitude" v:"required"`
	Longitude float64 `json:"longitude" v:"required"`
}

func (s *HubService) createHub(c *gin.Context) {
	var params createDto
	c.ShouldBindJSON(&params)

	if isValid, errMsg := validator.Validate(params); !isValid {
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return
	}

	hubData := hub.Hub{
		Name:      params.Name,
		City:      params.City,
		Area:      params.Area,
		Region:    params.Region,
		Country:   params.Country,
		Latitude:  params.Latitude,
		Longitude: params.Longitude,
	}

	result := db.Create(&hubData)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(http.StatusCreated, hubData)
}

func (s *HubService) getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	hub, err := s.Hub.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Not found",
			"status": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *HubService) getMany(c *gin.Context) {
	hubs, err := s.Hub.GetItems()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Not found",
			"status": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, hubs)
}
