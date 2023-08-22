package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"ygodraft/backend/model"
)

type ygoJwtAuthClient struct {
	secretKey []byte
}

// NewYgoJwtAuthClient creates a new YGOJwtAuthClient with the given secret key.
func NewYgoJwtAuthClient(secretKey string) model.YGOJwtAuthClient {
	return &ygoJwtAuthClient{
		secretKey: []byte(secretKey),
	}
}

func (jc *ygoJwtAuthClient) GenerateToken(email string, permission int) (string, error) {
	claims := model.YgoClaims{
		Email:      email,
		Permission: permission,
		RegisteredClaims: jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jc.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create signed string")
	}

	return signedToken, nil
}

func (jc *ygoJwtAuthClient) ValidateToken(tokenString string) (*model.YgoClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.YgoClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jc.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse with claims")
	}

	if claims, ok := token.Claims.(*model.YgoClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("the given token does not contain valid ygo claims")
}
