package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if an error occurred during processing
		if len(c.Errors) > 0 {
			// Get the last error that occurred
			err := c.Errors.Last().Err

			// Return an error response to the client
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}
}
