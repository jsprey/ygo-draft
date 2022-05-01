package ygoprodeck_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	ygoprodeck2 "ygodraft/backend/client/ygoprodeck"
)

func TestNewYgoProDeckClient(t *testing.T) {
	t.Run("create new client", func(t *testing.T) {
		// when
		client := ygoprodeck2.NewYgoProDeckClient()

		// then
		assert.NotNil(t, client)
		assert.NotNil(t, client.Client)
	})
}
