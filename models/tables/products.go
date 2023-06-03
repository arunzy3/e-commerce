package tables

type Products struct {
	ID       int     `gorm:"column:product_id;primaryKey;autoIncrement"`
	PID      string  `gorm:"column:product_pid;unique;not null"`
	Name     string  `gorm:"column:name"`
	Category string  `gorm:"column:category"`
	Quantity int     `gorm:"column:quantity"`
	Price    float64 `gorm:"column:price"`
}
