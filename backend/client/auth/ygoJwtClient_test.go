package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestYgoJwtAuthClient_NewYgoJwtAuthClient(t *testing.T) {
	// given
	secretKey := "my-secret-key"

	// when
	NewYgoJwtAuthClient(secretKey)
}

func TestYgoJwtAuthClient_GenerateToken(t *testing.T) {
	// given
	secretKey := "my-secret-key"
	client := NewYgoJwtAuthClient(secretKey)

	t.Run("runs correctly and returns a valid token", func(t *testing.T) {
		// given
		email := "test@example.com"
		permission := 1

		// when
		token, err := client.GenerateToken(email, permission)

		// then
		if err != nil {
			t.Errorf("GenerateToken failed: %v", err)
		}
		if len(token) == 0 {
			t.Errorf("GenerateToken produced an empty token")
		}
	})
}

func TestYgoJwtAuthClient_ValidateToken(t *testing.T) {
	// given
	secretKey := "my-secret-key"
	client := NewYgoJwtAuthClient(secretKey)

	t.Run("fails when parsing", func(t *testing.T) {
		// given
		invalidToken := "invalid-token"

		// when
		_, err := client.ValidateToken(invalidToken)

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse with claims")
	})

	t.Run("runs perfectly", func(t *testing.T) {
		// given
		email := "test@example.com"
		permission := 1

		token, err := client.GenerateToken(email, permission)
		if err != nil {
			t.Errorf("Failed to generate token for testing: %v", err)
			return
		}

		// when
		claims, err := client.ValidateToken(token)

		// then
		if err != nil {
			t.Errorf("ValidateToken failed: %v", err)
		}
		if claims.Email != email || claims.Permission != permission {
			t.Errorf("Validated token claims do not match")
		}
	})
}
