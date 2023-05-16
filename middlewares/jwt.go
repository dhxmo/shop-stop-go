package middlewares

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, utils.Response(nil, "Unauthorized", "30001"))
			c.Abort()
			return
		}

		data, err := utils.ValidateToken(token)
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, utils.Response(nil, "Unauthorized", "30001"))
			} else {
				c.JSON(http.StatusInternalServerError, utils.Response(nil, "Internal server error", "50000"))
			}
			c.Abort()
			return
		}

		c.Set("user", data)
		c.Next()
	}
}
