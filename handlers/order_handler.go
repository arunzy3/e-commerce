package handlers

import (
	"e-commerce/app"
	"e-commerce/services/order"
	"e-commerce/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var OrderStruct order.OrdersInterface = order.OrderImplementations{
	DbConnect: app.Connection(),
}

func PlaceOrderHandler(c *gin.Context) {
	db := OrderStruct.(order.OrderImplementations).DbConnect
	request, products, err := validatePlaceOrderHandler(c, db)
	if err.Error != "" {
		code := http.StatusUnprocessableEntity
		if err.Type == "internal_server_error" {
			code = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(code, err)
		return
	}
	orders, err := OrderStruct.PlaceOrder(request, products)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response := responsePlaceOrderCreator(orders, request)
	utils.ReturnJsonStruct(c, response)
}

func CancelOrderHandler(c *gin.Context) {
	pid, err := validateCancelOrderHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = OrderStruct.CancelOrder(pid)
	if err.Error != "" {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error, "unable to process the request. please try again after some time.") {
			statusCode = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(statusCode, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "order canceled succesfully",
	})
}

func ReadOrderHandler(c *gin.Context) {
	pid, err := validateReadOrderHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	product, productDetails, err := OrderStruct.ReadOrder(pid)
	if err.Error != "" {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error, "unable to process the request. please try again after some time.") {
			statusCode = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(statusCode, err)
		return
	}
	response := readOrderResponseConverter(product, productDetails)
	utils.ReturnJsonStruct(c, response)
}

func UpdateStatusHandler(c *gin.Context) {
	pid, err := validateUpdateStatusHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	product, productDetails, err := OrderStruct.UpdateStatus(pid)
	if err.Error != "" {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error, "unable to process the request. please try again after some time.") {
			statusCode = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(statusCode, err)
		return
	}
	response := readOrderResponseConverter(product, productDetails)
	utils.ReturnJsonStruct(c, response)
}
