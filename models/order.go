package models

type CreateOrderRequest struct {
	Orders []Orders `json:"orders"`
}

type Orders struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetOrderByIDResponse struct {
	ID             string   `json:"id"`
	ProductDetails []Orders `json:"product_details"`
	Status         string   `json:"status"`
	DispatchDate   string   `json:"dispatch_date,omitempty"`
	OrderValue     float64  `json:"order_value"`
	PayableAmount  float64  `json:"payable_amount"`
}
