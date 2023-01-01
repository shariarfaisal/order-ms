package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/brand"
	"github.com/shariarfaisal/order-ms/category"
	"github.com/shariarfaisal/order-ms/hub"
	"github.com/shariarfaisal/order-ms/order"
	"github.com/shariarfaisal/order-ms/product"
)


func initRoutes(router *gin.Engine) {
	hub.Routes(db, router)
	category.Routes(db, router)
	brand.Routes(db, router)
	product.Routes(db, router)
	order.Routes(db, router)
}