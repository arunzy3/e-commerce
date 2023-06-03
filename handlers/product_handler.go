package handlers

import (
	"e-commerce/app"
	"e-commerce/services/product"
	"e-commerce/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var ProductStruct product.ProductInterface = product.ProductImplementations{
	DbConnect: app.Connection(),
}

func AddProductHandler(c *gin.Context) {
	request, err := validateAddProductHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	product, err := ProductStruct.AddProduct(request)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	response := addProductResponseConverter(product)
	utils.ReturnJsonStruct(c, response)
}

func ReadProductHandler(c *gin.Context) {
	pid, err := validateReadProductHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	product, err := ProductStruct.ReadProduct(pid)
	if err.Error != "" {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error, "unable to process the request. please try again after some time.") {
			statusCode = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(statusCode, err)
		return
	}
	response := readProductResponseConverter(product)
	utils.ReturnJsonStruct(c, response)
}

func ReadProductsHandler(c *gin.Context) {
	category, isCategorySpecified, err := validateReadProductsHandler(c)
	if err.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	product, err := ProductStruct.ReadProducts(category, isCategorySpecified)
	if err.Error != "" {
		statusCode := http.StatusBadRequest
		if strings.Contains(err.Error, "unable to process the request. please try again after some time") {
			statusCode = http.StatusInternalServerError
		}
		c.AbortWithStatusJSON(statusCode, err)
		return
	}
	response := readProductsResponseConverter(product)
	utils.ReturnJsonStruct(c, response)
}
