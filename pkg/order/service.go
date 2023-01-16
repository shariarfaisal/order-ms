package order

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"gorm.io/gorm"
)

var db *gorm.DB

/*
Tables:
	Order
	OrderTimeline
	OrderCharges
	PickupStatus
	Pickup
	OrderItem
	DeliveryAddress
	AssignedRider
	CartItem
	Cart
	OrderIssue
	RestaurantPenalty
	RiderPenalty
	OrderRefund
*/

func Init(database *gorm.DB, r *gin.Engine) {
	db = database
	Migration(db)

	ordersGroup := r.Group("/orders")
	{
		ordersGroup.POST("/create", createOrder)
		// ordersGroup.GET("/:id", getOrder)
		// ordersGroup.GET("/", getOrders)
		// ordersGroup.PUT("/:id", updateOrder)
		// ordersGroup.DELETE("/:id", deleteOrder)
		// ordersGroup.GET("/:id/timeline", getOrderTimeline)
		// ordersGroup.GET("/:id/pickups", getOrderPickups)
		// ordersGroup.PUT("/:orderId/pickups/:id", orderPickupUpdate)
		// ordersGroup.DELETE("/:orderId/pickups/:id", orderPickupDelete)
		// ordersGroup.POST("/:id/issues", createOrderIssue)
		// ordersGroup.GET("/:id/issues", getOrderIssues)
		// ordersGroup.GET("/:orderId/issues/:id", getOrderIssueById)
	}

	//TODO: cartGroup := r.Group("/carts")
	{
		// cartGroup.GET("/", getCart) // get cart by user id, id take from auth token
		// cartGroup.POST("/item/create", addCartItem)
		// cartGroup.PUT("/item/:id", updateCartItem)
		// cartGroup.DELETE("/item/:id", deleteCartItem)
		// cartGroup.DELETE("/clear", clearCart)
	}
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
