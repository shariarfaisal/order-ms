package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/customer"
	"github.com/shariarfaisal/order-ms/hub"
	"github.com/shariarfaisal/order-ms/order"
	"github.com/shariarfaisal/order-ms/product"
	"github.com/shariarfaisal/order-ms/rider"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=orderms port=5432 sslmode=disable"
	dbRes, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = dbRes
	order.Migration(db)
	hub.Migration(db)
	rider.Migration(db)
	brand.Migration(db)
	product.Migration(db)
	customer.Migration(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Use(JSONMiddleware())

	initRoutes(r)

	r.Run(":5000") // listen and serve on

}
