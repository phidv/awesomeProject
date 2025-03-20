package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oms/pkg/constants"
	"oms/pkg/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.WriteErrorResponse(c, fmt.Errorf("%s", constants.ErrMissingAuthHeader), http.StatusUnauthorized)
			return
		}

		//"Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteErrorResponse(c, fmt.Errorf("%s", constants.ErrInvalidTokenFormat), http.StatusUnauthorized)
			return
		}

		claims, err := utils.VerifyToken(parts[1])
		if err != nil {
			utils.WriteErrorResponse(c, fmt.Errorf("%s", constants.ErrInvalidToken), http.StatusUnauthorized)
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
