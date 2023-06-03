package utils

import (
	"bytes"
	"e-commerce/models"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func ReturnJsonStruct(c *gin.Context, genericStruct interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Set("responsestring", fmt.Sprint(genericStruct))
	json.NewEncoder(c.Writer).Encode(genericStruct)
}

func ValidateUnknownParams(reqBody interface{}, ctx *gin.Context) models.Errors {
	ginErr := models.Errors{
		Error: "the request body is invalid",
		Type:  "invalid_request_error",
	}
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqBody)
	if err != nil {
		return ginErr
	}
	payloadBS, err := json.Marshal(&reqBody)
	if err != nil {
		return ginErr
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(payloadBS))
	return models.Errors{}
}
