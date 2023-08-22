package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
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
		user := model.User{
			Email:       "test@example.com",
			DisplayName: "Test User",
			IsAdmin:     true,
		}

		// when
		token, err := client.GenerateToken(user)

		// then
		require.NoError(t, err)
		assert.NotNil(t, token)
		assert.True(t, len(token) > 0)
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
		user := model.User{
			Email:       "test@example.com",
			DisplayName: "Test User",
			IsAdmin:     true,
		}

		token, err := client.GenerateToken(user)
		if err != nil {
			t.Errorf("Failed to generate token for testing: %v", err)
			return
		}

		// when
		claims, err := client.ValidateToken(token)

		// then
		require.NoError(t, err)
		assert.Equal(t, user.Email, claims.Email)
		assert.Equal(t, user.IsAdmin, claims.IsAdmin)
		assert.Equal(t, user.DisplayName, claims.DisplayName)
	})
}
