package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/customer"
	"github.com/shariarfaisal/validator"
	"gorm.io/gorm"
)

type OrderItemSchema struct  {
	Id int `json:"id" v:"required"`
	Quantity int `json:"quantity" v:"required;min=1;max=100"`
}

type OrderSchema struct {
	Platform string `json:"platform" v:"required;enum=web,app"`
	RiderNote string `json:"riderNote"`
	Items []OrderItemSchema`json:"items" v:"required;min=1;"`
	RestaurantNote string `json:"restaurantNote"`
	AddressId int `json:"addressId" v:"required"`
	PaymentMethod string `json:"paymentMethod" v:"required"`
	Voucher string `json:"voucher" v:"required"`
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


func createOrder(c *gin.Context) {
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

	_, err := customer.GetCustomerById(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = customer.GetAddressById(params.AddressId)
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

	db.Transaction(func (tx *gorm.DB) error {

		// create order 
		// create pickups 
		// create order items

		return nil
	})

	c.JSON(http.StatusOK, gin.H{"message": "Order Created"})
}

