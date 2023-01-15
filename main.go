package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	dotEnv "github.com/joho/godotenv"
	"github.com/shariarfaisal/order-ms/pkg/admin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/customer"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"github.com/shariarfaisal/order-ms/pkg/order"
	"github.com/shariarfaisal/order-ms/pkg/product"
	"github.com/shariarfaisal/order-ms/pkg/rider"
	"github.com/shariarfaisal/order-ms/pkg/section"
	"github.com/shariarfaisal/order-ms/pkg/voucher"
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

func initEnv() {
	err := dotEnv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	initEnv()
	env := os.Getenv("ENV")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := "disable"

	if env == "production" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)
	dbRes, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = dbRes

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	admin.Init(db, r)
	hub.Init(db, r)
	order.Init(db, r)
	rider.Init(db, r)
	product.Init(db, r)
	customer.Init(db, r)
	voucher.Init(db, r)
	section.Init(db, r)
	brand.Init(db, r)

	r.Use(JSONMiddleware())

	r.Run(":5000") // listen and serve on

}
