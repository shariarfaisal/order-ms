package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/pkg/brand"
	"github.com/shariarfaisal/order-ms/pkg/category"
	"github.com/shariarfaisal/order-ms/pkg/hub"
	"github.com/shariarfaisal/order-ms/pkg/order"
	"github.com/shariarfaisal/order-ms/pkg/product"
)

func initRoutes(router *gin.Engine) {
	hub.Routes(db, router)
	category.Routes(db, router)
	brand.Routes(db, router)
	product.Routes(db, router)
	order.Routes(db, router)
}
