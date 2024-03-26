package auth

import (
	"fmt"
	"net/http"
	"strings"
	"ygodraft/backend/model"

	"github.com/gin-gonic/gin"
)

const ContextClaimKey = "user_claims"

func PermissionMiddleware(client model.YGOJwtAuthClient, isAdminPermission bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			err := fmt.Errorf("no authorization header found")
			ctx.String(http.StatusUnauthorized, err.Error())
			_ = ctx.AbortWithError(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(tokenHeader, " ")[1] // Assuming token format is "Bearer token"
		tokenClaims, err := client.ValidateToken(tokenString)
		if err != nil {
			err := fmt.Errorf("given token not valid")
			ctx.String(http.StatusUnauthorized, err.Error())
			_ = ctx.AbortWithError(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		ctx.Set(ContextClaimKey, tokenClaims)

		// Check if the user has the required permission
		if isAdminPermission && !tokenClaims.IsAdmin {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// GetClaims retrieves the user claims from the context if they are available, if not false is returned.
func GetClaims(ctx *gin.Context) (*model.YgoClaims, bool) {
	claims, exists := ctx.Get(ContextClaimKey)
	if !exists {
		return nil, false
	}

	ygoClaims, ok := claims.(*model.YgoClaims)
	if !ok {
		return nil, false
	}

	return ygoClaims, true
}
