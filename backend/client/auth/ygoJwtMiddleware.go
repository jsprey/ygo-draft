package auth

import (
	"fmt"
	"net/http"
	"strings"
	"ygodraft/backend/model"

	"github.com/gin-gonic/gin"
)

const claimNamePermission = "permission"

// PermissionAdmin is the permission for admin users
const PermissionAdmin int = 100

// PermissionUser is the permission for normal users
const PermissionUser int = 10

func PermissionMiddleware(client model.YGOJwtAuthClient, targetPermission int) gin.HandlerFunc {
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
		if tokenClaims.Permission < targetPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
