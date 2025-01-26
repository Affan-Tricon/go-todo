package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data, meta interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Ok",
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

func BadResponse(c *gin.Context, message string, data, meta interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "Error",
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

func UnAuthorizeResponse(c *gin.Context, message string, data, meta interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "Error",
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

func InternalErrorResponse(c *gin.Context, message string, data, meta interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "Error",
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}
