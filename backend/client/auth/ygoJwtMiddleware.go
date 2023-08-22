package auth

import (
	"fmt"
	"net/http"
	"strings"
	"ygodraft/backend/model"

	"github.com/gin-gonic/gin"
)

const claimNamePermission = "permission"

func PermissionMiddleware(client model.YGOJwtAuthClient, isAdminPermission bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			err := fmt.Errorf("no authorization header found")
			c.String(http.StatusUnauthorized, err.Error())
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		tokenString := strings.Split(tokenHeader, " ")[1] // Assuming token format is "Bearer token"
		tokenClaims, err := client.ValidateToken(tokenString)
		if err != nil {
			err := fmt.Errorf("given token not valid")
			c.String(http.StatusUnauthorized, err.Error())
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		// Check if the user has the required permission
		if isAdminPermission && !tokenClaims.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
