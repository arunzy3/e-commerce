package product

import (
	"e-commerce/models"
	"e-commerce/models/tables"
	"e-commerce/utils"
	"errors"

	"gorm.io/gorm"
)

type ProductImplementations struct {
	DbConnect *gorm.DB
}

func (p ProductImplementations) ReadProduct(pid string) (tables.Products, models.Errors) {
	var product tables.Products
	err := p.DbConnect.Where("product_pid = ?", pid).Take(&product).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return product, models.Errors{
				Error: "unable to process the request. please try again after some time",
				Type:  "internal_request_error",
			}
		}
		return product, models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "product_id",
		}
	}
	return product, models.Errors{}
}

func (p ProductImplementations) ReadProducts(category string, isCategorySpecified bool) ([]tables.Products, models.Errors) {
	var products []tables.Products
	var err error
	if isCategorySpecified {
		err = p.DbConnect.Where("category = ?", category).Find(&products).Error
	} else {
		err = p.DbConnect.Find(&products).Error
	}
	if err != nil {
		return products, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_request_error",
		}
	}
	if isCategorySpecified && len(products) == 0 {
		return products, models.Errors{
			Error: "there is no product with the mentioned category",
			Type:  "invalid_request_error",
			Param: "category",
		}
	}
	return products, models.Errors{}
}

func (p ProductImplementations) AddProduct(reqBody models.AddProductRequest) (tables.Products, models.Errors) {
	product := tables.Products{
		PID:      utils.GenerateUUID("pr"),
		Name:     reqBody.Name,
		Category: reqBody.Category,
		Quantity: reqBody.Quantity,
		Price:    reqBody.Price,
	}
	err := p.DbConnect.Create(&product).Error
	if err != nil {
		return product, models.Errors{
			Error: "unable to process the request. please try again after some time",
			Type:  "internal_request_error",
		}
	}
	return product, models.Errors{}
}
