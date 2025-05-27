package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 是API响应的标准格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

// SuccessResponse 返回成功响应
func SuccessResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// ErrorResponse 返回错误响应
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
	})
}

// BadRequestResponse 返回请求错误响应
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

// NotFoundResponse 返回资源不存在响应
func NotFoundResponse(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	ErrorResponse(c, http.StatusNotFound, message)
}

// InternalServerErrorResponse 返回服务器内部错误响应
func InternalServerErrorResponse(c *gin.Context, message string) {
	if message == "" {
		message = "Internal server error"
	}
	ErrorResponse(c, http.StatusInternalServerError, message)
}
