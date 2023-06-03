package handlers

import (
	"e-commerce/models"
	"e-commerce/models/tables"
	"e-commerce/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func validatePlaceOrderHandler(c *gin.Context, db *gorm.DB) (models.CreateOrderRequest, []tables.Products, models.Errors) {
	var request models.CreateOrderRequest
	var products []tables.Products
	err := utils.ValidateUnknownParams(&request, c)
	if err.Error != "" {
		return request, products, err
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		return request, products, models.Errors{
			Error: "the request body is invalid",
			Type:  "invalid_request_error",
		}
	}
	if len(request.Orders) == 0 {
		return request, products, models.Errors{
			Error: "the request body is invalid",
			Type:  "invalid_request_error",
		}
	}
	mappingOfQuantity := make(map[string]int)
	for _, order := range request.Orders {
		if order == (models.Orders{}) ||
			((order.ProductID == "") && (order.Quantity != 0)) ||
			((order.ProductID != "") && (order.Quantity == 0)) {
			return request, products, models.Errors{
				Error: "the request body is invalid",
				Type:  "invalid_request_error",
			}
		}
		var product tables.Products
		dbErr := db.Where("product_pid = ?", order.ProductID).Take(&product).Error
		if dbErr != nil {
			if !errors.Is(dbErr, gorm.ErrRecordNotFound) {
				return request, products, models.Errors{
					Error: "unable to process the request. please try again after some time",
					Type:  "internal_server_error",
				}
			}
			return request, products, models.Errors{
				Error: "invalid product_id in the list : " + order.ProductID,
				Type:  "invalid_request_error",
				Param: "product_id",
			}
		}
		if order.Quantity > 10 {
			return request, products, models.Errors{
				Error: "the quantity cannot exceed 10",
				Type:  "invalid_request_error",
				Param: "quantity",
			}
		}
		value, ok := mappingOfQuantity[order.ProductID]
		if ok {
			mappingOfQuantity[order.ProductID] = value + order.Quantity
		} else {
			mappingOfQuantity[order.ProductID] = order.Quantity
		}
		products = append(products, product)
	}
	for key, value := range mappingOfQuantity {
		var product tables.Products
		for _, singleProduct := range products {
			if key == singleProduct.PID {
				product = singleProduct
				break
			}
		}
		if value > product.Quantity {
			return request, products, models.Errors{
				Error: "not enough stock",
				Type:  "invalid_request_error",
			}
		}
	}
	return request, products, models.Errors{}
}

func responsePlaceOrderCreator(orders tables.Orders, req models.CreateOrderRequest) models.GetOrderByIDResponse {
	return models.GetOrderByIDResponse{
		ID:             orders.PID,
		ProductDetails: req.Orders,
		Status:         orders.Status,
		OrderValue:     orders.OrderValue,
		PayableAmount:  orders.PayableAmount,
	}
}

func validateCancelOrderHandler(c *gin.Context) (string, models.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, models.Errors{}
}

func validateReadOrderHandler(c *gin.Context) (string, models.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, models.Errors{}
}

func readOrderResponseConverter(order tables.Orders, productDetails []models.Orders) (response models.GetOrderByIDResponse) {
	return models.GetOrderByIDResponse{
		ID:             order.PID,
		ProductDetails: productDetails,
		Status:         order.Status,
		DispatchDate:   order.DispatchDate,
		OrderValue:     order.OrderValue,
		PayableAmount:  order.PayableAmount,
	}
}

func validateUpdateStatusHandler(c *gin.Context) (string, models.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	return pid, models.Errors{}
}
