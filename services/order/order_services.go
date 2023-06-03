package order

import (
	"e-commerce/models"
	"e-commerce/models/tables"
	"e-commerce/utils"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type OrderImplementations struct {
	DbConnect *gorm.DB
}

func (o OrderImplementations) PlaceOrder(reqBody models.CreateOrderRequest, products []tables.Products) (tables.Orders, models.Errors) {
	var orderValue float64 = 0
	var payableAmount float64 = 0
	mapping := make(map[string]int)
	for _, order := range reqBody.Orders {
		var productToBeUpdated tables.Products

		for _, product := range products {
			if order.ProductID == product.PID {
				productToBeUpdated = product
				break
			}
		}
		if productToBeUpdated.Category == "premium" {
			mapping[order.ProductID]++
		}
		productToBeUpdated.Quantity = productToBeUpdated.Quantity - order.Quantity
		err := o.DbConnect.Save(&productToBeUpdated).Error
		if err != nil {
			return tables.Orders{}, models.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		}
		orderValue = orderValue + (float64(order.Quantity) * productToBeUpdated.Price)
	}

	payableAmount = orderValue
	if len(mapping) > 2 {
		payableAmount = orderValue * 0.9
	}
	productDetailsJSON, err := json.Marshal(reqBody.Orders)
	if err != nil {
		return tables.Orders{}, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	newOrder := tables.Orders{
		PID:            utils.GenerateUUID("or"),
		Status:         "placed",
		OrderValue:     orderValue,
		PayableAmount:  payableAmount,
		ProductDetails: string(productDetailsJSON),
	}
	err = o.DbConnect.Create(&newOrder).Error
	if err != nil {
		return tables.Orders{}, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	return newOrder, models.Errors{}
}

func (o OrderImplementations) CancelOrder(pid string) models.Errors {
	var order tables.Orders
	err := o.DbConnect.Where("order_pid = ?", pid).Take(&order).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		}
		return models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	if order.Status == "completed" ||
		order.Status == "canceled" {
		return models.Errors{
			Error: "the order has already been " + order.Status,
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	err = o.DbConnect.Model(&tables.Orders{}).Where("order_pid = ?", pid).UpdateColumn("status", "canceled").Error
	if err != nil {
		return models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	return models.Errors{}
}

func (o OrderImplementations) ReadOrder(pid string) (tables.Orders, []models.Orders, models.Errors) {
	var order tables.Orders
	err := o.DbConnect.Where("order_pid = ?", pid).Take(&order).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return order, []models.Orders{}, models.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		}
		return order, []models.Orders{}, models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}
	var productDetails []models.Orders
	json.Unmarshal([]byte(order.ProductDetails), &productDetails)
	if err != nil {
		return order, []models.Orders{}, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	return order, productDetails, models.Errors{}
}

func (o OrderImplementations) UpdateStatus(pid string) (tables.Orders, []models.Orders, models.Errors) {
	var order tables.Orders
	err := o.DbConnect.Where("order_pid = ?", pid).Take(&order).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return order, []models.Orders{}, models.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_server_error",
			}
		}
		return order, []models.Orders{}, models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}

	if order.Status != "dispatched" &&
		order.Status != "placed" {
		return order, []models.Orders{}, models.Errors{
			Error: "the order has already been " + order.Status,
			Type:  "invalid_request_error",
			Param: "order_id",
		}
	}

	switch order.Status {
	case "placed":
		order.Status = "dispatched"
		order.DispatchDate = time.Now().Format("2006-01-02")
	case "dispatched":
		order.Status = "completed"
	}

	err = o.DbConnect.Save(&order).Error
	if err != nil {
		return order, []models.Orders{}, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	var productDetails []models.Orders
	json.Unmarshal([]byte(order.ProductDetails), &productDetails)
	if err != nil {
		return order, []models.Orders{}, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_server_error",
		}
	}
	return order, productDetails, models.Errors{}
}
