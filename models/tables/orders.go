package tables

type Orders struct {
	ID             int     `gorm:"column:order_id;primaryKey;autoIncrement"`
	PID            string  `gorm:"column:order_pid;unique;not null"`
	Status         string  `gorm:"column:status"`
	DispatchDate   string  `gorm:"column:dispatch_date"`
	OrderValue     float64 `gorm:"column:order_value"`
	PayableAmount  float64 `gorm:"column:payable_amount"`
	ProductDetails string  `gorm:"column:product_details;type:json"`
}
