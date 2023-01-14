package order

import (
	"github.com/gin-gonic/gin"
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

func Routes(database *gorm.DB, r *gin.Engine) {
	db = database

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
