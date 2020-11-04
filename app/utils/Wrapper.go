package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WrapAPIError untuk menampilkan pesan error
func WrapAPIError(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"Code":          code,
		"Error_Type":    http.StatusText(code),
		"Error_Details": message,
	})
}

// WrapAPISuccess untuk menampilkan pesan sukses
func WrapAPISuccess(c *gin.Context, message string, code int) {
	c.JSON(code, map[string]interface{}{
		"Code":   code,
		"Status": message,
	})
}

// WrapAPIData untuk menampilkan data
func WrapAPIData(c *gin.Context, data interface{}, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"Code":   code,
		"Status": message,
		"Data":   data,
	})
}
