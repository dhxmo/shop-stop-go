package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code string

		code = "200"
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.Response(nil, "Unauthorized", "400"))

			c.Abort()
			return
		}

		_, err := utils.ValidateToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = "30002"
			default:
				code = "30001"
			}
		}

		if code != "200" {
			c.JSON(http.StatusUnauthorized, utils.Response(nil, "Unauthorized", code))

			c.Abort()
			return
		}

		c.Next()
	}
}
