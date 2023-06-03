package app

import (
	"e-commerce/models/tables"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&tables.Orders{},
		&tables.Products{})
	if err != nil {
		panic("Database Migration failed")
	}
}
