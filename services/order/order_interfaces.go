package order

import (
	"e-commerce/models"
	"e-commerce/models/tables"
)

type OrdersInterface interface {
	PlaceOrder(reqBody models.CreateOrderRequest, products []tables.Products) (tables.Orders, models.Errors)
	ReadOrder(pid string) (tables.Orders, []models.Orders, models.Errors)
	UpdateStatus(pid string) (tables.Orders, []models.Orders, models.Errors)
	CancelOrder(pid string) models.Errors
}
