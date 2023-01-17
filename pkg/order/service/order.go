package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/market"
	"github.com/shariarfaisal/order-ms/pkg/order"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type OrderService struct {
	Order    *order.OrderRepo
	Customer *market.CustomerRepo
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		Order:    order.NewOrderRepo(db),
		Customer: market.NewCustomerRepo(db),
	}
}

type OrderItemSchema struct {
	Id       int `json:"id" v:"required"`
	Quantity int `json:"quantity" v:"required;min=1;max=100"`
}

type OrderSchema struct {
	Platform       string            `json:"platform" v:"required;enum=web,app"`
	RiderNote      string            `json:"riderNote"`
	Items          []OrderItemSchema `json:"items" v:"required;min=1;"`
	RestaurantNote string            `json:"restaurantNote"`
	AddressId      int               `json:"addressId" v:"required"`
	PaymentMethod  string            `json:"paymentMethod" v:"required"`
	Voucher        string            `json:"voucher" v:"required"`
}

/**
fetch products & products restaurant & Hub by items ids
	- check products existence
	- check products status
	- check products inventory
	- check restaurant operating status
	TODO check group restaurant validity
	- Check hub order accepting or not
	- Check is order from multiple hub or not (validate)
	- Calculate order items total price & discount
- Voucher Validation & calculate discount
- Get delivery charge
- Check payment Status

TODO: Create Transaction
	- Create Order
	- Create Pickups
	- Create Order Items
	- Update inventory
	- Update voucher info
	- Update Brands counter
	- Update Products counter

TODO: Send Notification
	- Notify customer
	- Notify restaurant
	- Notify Operation
*/

func (s *OrderService) createOrder(c *gin.Context) {
	var params OrderSchema
	c.ShouldBindJSON(&params)

	isValid, errors := validator.Validate(params)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	/**
	- Input data validation (required, type, length, etc)
	- Check customer validity (is customer exists, is customer active, etc)
	- Check address validity (is address exists, is address valid etc)
	- check operation status
	*/

	_, err := s.Customer.GetById(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = s.Customer.GetAddressById(params.AddressId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems, err := getOrderItems(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brands := []brand.Brand{}
	for _, item := range orderItems {
		brands = append(brands, item.Brand)
	}

	isOperating, errMsg := isBrandsOperating(brands)
	if !isOperating {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	// hubId := order.HubId
	// _, err := hub.HubById(hubId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// orderObj := Order {
	// 	Status: order.Status,
	// 	Platform: order.Platform,
	// 	DispatchTime: order.DispatchTime,
	// 	RiderNote: order.RiderNote,
	// 	HubId: order.HubId,
	// 	EDT: order.EDT,
	// 	PaymentMethod: order.PaymentMethod,
	// 	PaymentStatus: order.PaymentStatus,
	// 	Total: order.Total,
	// 	Discount: order.Discount,
	// 	ServiceCharge: order.ServiceCharge,
	// 	ItemsValue: order.ItemsValue,
	// 	ItemDiscount: order.ItemDiscount,
	// 	PromoDiscount: order.PromoDiscount,
	// }

	db.Transaction(func(tx *gorm.DB) error {

		// create order
		// create pickups
		// create order items

		return nil
	})

	c.JSON(http.StatusOK, gin.H{"message": "Order Created"})
}

func isValidItemsForOrder(items []OrderItemSchema, prods []brand.Product) (bool, string) {
	/*
		- check products existence
		- check products status
		- check products inventory
	*/

	byId := map[int]brand.Product{}
	for _, prod := range prods {
		byId[int(prod.ID)] = prod
	}

	for _, item := range items {
		prod, ok := byId[item.Id]
		if !ok || prod.Status != brand.ProductStatusApproved {
			return false, fmt.Sprintf("%s not found", prod.Name)
		}

		if !prod.IsAvailable {
			return false, fmt.Sprintf("%s not available", prod.Name)
		}

		if prod.UseInventory && prod.Stock < item.Quantity {
			return false, fmt.Sprintf("%s out of stock", prod.Name)
		}
	}

	return true, ""
}

func getOrderItems(params OrderSchema) ([]brand.Product, error) {
	productRepo := brand.NewProductRepo(db)

	items := params.Items
	itemsId := []int{}
	for _, item := range items {
		itemsId = append(itemsId, item.Id)
	}

	prods, err := productRepo.GetByIds(itemsId)
	if err != nil {
		return nil, err
	}

	_, errMsg := isValidItemsForOrder(items, prods)
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	return prods, nil
}

func InBetweenHours(start /* start hour */, end /* end hour*/, hour, minute float32) bool {
	if minute > 0 {
		hour += minute / 100
	}

	if start < end {
		return start <= hour && end >= hour
	} else {
		return start <= hour || end >= hour
	}
}

func isBrandsOperating(brands []brand.Brand) (bool, string) {

	for _, item := range brands {
		if item.Status != brand.BrandStatusActive {
			return false, fmt.Sprintf("%s is not operating", item.Name)
		}

		if !item.IsAvailable {
			return false, fmt.Sprintf("%s is not operating", item.Name)
		}

		opTime := item.OperatingTimes

		if len(opTime) > 0 {
			day := time.Now().Day()

			times, exists := opTime[strconv.Itoa(day)]
			if !exists {
				return false, fmt.Sprintf("%s is not operating", item.Name)
			}

			hour := float32(time.Now().Hour())
			minute := float32(time.Now().Minute())
			isOperating := false
			for _, t := range times.([]brand.OperatingTime) {
				from := t.From.Hour + t.From.Minute/100
				to := t.To.Hour + t.To.Minute/100
				if InBetweenHours(float32(from), float32(to), hour, minute) {
					isOperating = true
					break
				}
			}

			if !isOperating {
				return false, fmt.Sprintf("%s is not operating", item.Name)
			}
		}
	}

	return true, ""
}
