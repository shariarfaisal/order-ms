package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shariarfaisal/order-ms/order"
)


func initRoutes(router *gin.Engine) {
	order.Routes(db, router)
}