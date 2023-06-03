package main

import (
	"e-commerce/app"
	"e-commerce/handlers"
	"e-commerce/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := app.Connection()
	app.AutoMigration(db)

	router.Use(middleware.PanicHandling())

	productAPIs := router.Group("/product")
	productAPIs.POST("/add", handlers.AddProductHandler)
	productAPIs.GET("/list-all", handlers.ReadProductsHandler)
	productAPIs.GET("/:id", handlers.ReadProductHandler)

	orderAPIs := router.Group("/order")
	orderAPIs.POST("/add", handlers.PlaceOrderHandler)
	orderAPIs.GET("/:id", handlers.ReadOrderHandler)
	orderAPIs.PATCH("/status/:id", handlers.UpdateStatusHandler)
	orderAPIs.PATCH("/cancel/:id", handlers.CancelOrderHandler)
	router.Run()
}
