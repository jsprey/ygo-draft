package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type YgoClaims struct {
	Email      string `json:"email"`
	Permission int    `json:"permission"`
	jwt.RegisteredClaims
}

type YGOJwtAuthClient interface {
	GenerateToken(username string, permission int) (string, error)
	ValidateToken(tokenString string) (*YgoClaims, error)
}
