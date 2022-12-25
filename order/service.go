package order

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/product"
	"gorm.io/gorm"
)

/**
Expected body data
{
	"platform": "web" | "app",
	"rider_note": "string",
	"items": [
		{
			"id": 1,
			"quantity": 1,
			"note": "string" // optional
		}
	],
	"restaurant_note": "string", // optional
	"address_id": number, // optional
	"payment_method": "cash" | "bkash" | "ssl" | "aamarpay", // optional
	"voucher": "string", // optional
}
*/ 

type DTO struct {
	Platform string `json:"platform" validate:"required"`
	RiderNote string `json:"rider_note"`
	Items [] struct {
		Id int `json:"id" validate:"required"`
		Quantity int `json:"quantity" validate:"required"`
		Note string `json:"note"`
	} `json:"items" validate:"required"`
	RestaurantNote string `json:"restaurant_note"`
	AddressId int `json:"address_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	Voucher string `json:"voucher" validate:"required"`
}

/*
	- Input data validation (required, type, length, etc)
	- Write error messages for specific field
*/

func Validate (d *DTO) {
	x := reflect.TypeOf(d)
	for i := 0; i < x.NumField(); i++ {
		println(x.Field(i).Tag.Get("validate"))
		// if x.Field(i).Tag.Get("validate") == "required" {

		// }
	}
}


func createOrder(c *gin.Context) {
	var params DTO
	c.BindJSON(&params)

	c.JSON(http.StatusOK, gin.H{"message": "Order Created", "data": params})

	// validate params
	

	/**
		- Input data validation (required, type, length, etc)
		- Check customer validity (is customer exists, is customer active, etc)
		- Check address validity (is address exists, is address valid etc)
		- check operation status 
	*/

	/**
		TODO: fetch products & products restaurant & Hub by items ids
			- check products existence 
			- check products status 
			- check products inventory 
			- check restaurant operating status
			- check group restaurant validity
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


	// hubId := order.HubId
	// _, err := hub.HubById(hubId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	itemsId := []int{}
	for _, item := range params.Items {
		println(item.Quantity)
		itemsId = append(itemsId, item.Id)
	}

	prods := product.GetByIds(itemsId)
	print(len(prods))


	

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

