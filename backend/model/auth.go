package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// YgoClaims contains the claims for the user contained in the jwt tokens.
type YgoClaims struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	IsAdmin     bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// YGOJwtAuthClient is used to handle the authentication of the jwt tokens.
type YGOJwtAuthClient interface {
	GenerateToken(user User) (string, error)
	ValidateToken(tokenString string) (*YgoClaims, error)
}
