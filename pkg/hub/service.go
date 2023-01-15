package hub

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)
	hr := r.Group("/hub")
	{
		hr.GET("/", getMany)
		hr.GET("/:id", getById)
		hr.POST("/create", createHub)
		// hr.PUT("/:id", updateHub)
		// hr.DELETE("/:id", deleteHub)
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

func createHub(c *gin.Context) {
	var params createDto
	c.ShouldBindJSON(&params)

	if isValid, errMsg := validator.Validate(params); !isValid {
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return
	}

	hubData := Hub{
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
	c.JSON(http.StatusCreated, gin.H{
		"result": hubData,
	})
}

func getById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	hub, err := GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Not found",
			"status": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hub,
	})
}

func getMany(c *gin.Context) {
	hubs, err := GetMany()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Not found",
			"status": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": hubs,
	})
}
