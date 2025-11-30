package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func Deleted(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
	})
}

func ValidationError(c *gin.Context, statusCode int, err error) {
	var errorMessages []string

	for _, fieldErr := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, fieldErr.Error())
	}

	c.JSON(statusCode, Response{
		Success: false,
		Message: "validation failed",
		Error:   errorMessages,
	})
}
