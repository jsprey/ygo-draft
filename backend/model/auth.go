package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type YgoClaims struct {
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	IsAdmin     bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type YGOJwtAuthClient interface {
	GenerateToken(user User) (string, error)
	ValidateToken(tokenString string) (*YgoClaims, error)
}
