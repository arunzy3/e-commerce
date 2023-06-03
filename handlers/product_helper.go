package handlers

import (
	"e-commerce/models"
	"e-commerce/models/tables"
	"e-commerce/utils"

	"github.com/gin-gonic/gin"
)

func validateAddProductHandler(c *gin.Context) (models.AddProductRequest, models.Errors) {
	var request models.AddProductRequest
	err := utils.ValidateUnknownParams(&request, c)
	if err.Error != "" {
		return request, err
	}
	if err := c.ShouldBindJSON(&request); err != nil || ((models.AddProductRequest{}) == request) {
		return request, models.Errors{
			Error: "the request body is invalid",
			Type:  "invalid_request_error",
		}
	}
	if request.Category != "premium" &&
		request.Category != "regular" &&
		request.Category != "budget" {
		return request, models.Errors{
			Error: "invalid category",
			Type:  "invalid_request_error",
			Param: "category",
		}
	}
	return request, models.Errors{}
}

func validateReadProductsHandler(c *gin.Context) (string, bool, models.Errors) {
	category, isCategorySpecified := c.GetQuery("category")
	if isCategorySpecified {
		if category != "premium" &&
			category != "regular" &&
			category != "budget" {
			return category, isCategorySpecified, models.Errors{
				Error: "invalid category",
				Type:  "invalid_request_error",
				Param: "category",
			}
		}
	}
	return category, isCategorySpecified, models.Errors{}
}

func validateReadProductHandler(c *gin.Context) (string, models.Errors) {
	pid := c.Param("id")
	if pid == "" {
		return "", models.Errors{
			Error: "invalid id",
			Type:  "invalid_request_error",
			Param: "product_id",
		}
	}
	return pid, models.Errors{}
}

func addProductResponseConverter(product tables.Products) models.AddProductResponse {
	return models.AddProductResponse{
		ID:       product.PID,
		Name:     product.Name,
		Category: product.Category,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
}

func readProductResponseConverter(product tables.Products) (response models.AddProductResponse) {
	return models.AddProductResponse{
		ID:       product.PID,
		Name:     product.Name,
		Category: product.Category,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
}

func readProductsResponseConverter(products []tables.Products) (response models.ReadProductResponse) {
	var productData []models.AddProductResponse
	for _, product := range products {
		responseTemp := models.AddProductResponse{
			ID:       product.PID,
			Name:     product.Name,
			Category: product.Category,
			Quantity: product.Quantity,
			Price:    product.Price,
		}
		productData = append(productData, responseTemp)
	}
	response.Products = productData
	return response
}
