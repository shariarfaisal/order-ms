package service

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/market"
	"github.com/shariarfaisal/order-ms/pkg/utils"
	"github.com/shariarfaisal/order-ms/pkg/validator"
	"gorm.io/gorm"
)

type CustomerService struct {
	Customer *market.CustomerRepo
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		Customer: market.NewCustomerRepo(db),
	}
}

func (s *CustomerService) getProfile(c *gin.Context) {
	payload, err := c.Get("customer")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	id, isId := payload.(map[string]interface{})["id"]
	if !isId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	customer, er := s.Customer.GetBy("id", id)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, customer)
}

type LoginSchema struct {
	Phone    string `json:"phone"`
	Password string `json:"password" v:"required;min=6;"`
	Email    string `json:"email"`
}

func (s *CustomerService) login(c *gin.Context) {
	var payloads LoginSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	if payloads.Phone == "" && payloads.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone or Email is required",
		})
		return
	}

	var customer market.Customer
	if payloads.Phone != "" {
		data, err := s.Customer.GetBy("phone", payloads.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid phone or password",
			})
			return
		}
		customer = data
	} else {
		data, err := s.Customer.GetBy("email", payloads.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		customer = data
	}

	if !utils.IsValidPassword(payloads.Password, customer.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid phone or password",
		})
		return
	}

	jwtPayload := map[string]interface{}{
		"id":    customer.ID,
		"name":  customer.Name,
		"phone": customer.Phone,
		"email": customer.Email,
	}

	token, err := utils.GenerateJWT("customer", jwtPayload, os.Getenv("APP_SECRET"), 60*5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type SignUpSchema struct {
	Name        string `json:"name" v:"required;min=3;"`
	Email       string `json:"email" v:"email"`
	Phone       string `json:"phone" v:"required;min=11;max=11;"`
	Password    string `json:"password" v:"required;min=6;"`
	Image       string `json:"image"`
	Gender      string `json:"gender" v:"enum=male|female"`
	DateOfBirth string `json:"dateOfBirth" v:"date"`
}

func (s *CustomerService) signUp(c *gin.Context) {
	var payloads SignUpSchema
	c.ShouldBindJSON(&payloads)

	isValid, errors := validator.Validate(payloads)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	/*
		is email exist
		is phone exist
		hash password
	*/

	if exist, err := s.Customer.IsExist("phone", payloads.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone already exist",
		})
		return
	}

	if exist, err := s.Customer.IsExist("email", payloads.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exist",
		})
		return
	}

	customer := market.Customer{
		Name:     payloads.Name,
		Email:    payloads.Email,
		Phone:    payloads.Phone,
		Password: payloads.Password,
		Image:    payloads.Image,
		Gender:   market.CustomerGender(payloads.Gender),
		Status:   market.CustomerStatusActive,
		IsActive: true,
	}

	if payloads.DateOfBirth != "" {
		dateOfBirth, err := utils.ParseDate(payloads.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		customer.DateOfBirth = dateOfBirth
	}

	hashPass, err := utils.HashPassword(customer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.Password = hashPass

	err = db.Create(&customer).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
	})
}
