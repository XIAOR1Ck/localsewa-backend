package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    any  `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, status int, message string, data any) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
