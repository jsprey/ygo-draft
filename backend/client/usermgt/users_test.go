package usermgt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model/mocks"
)

func TestNewUsermgtClient(t *testing.T) {
	// when
	client, err := NewUsermgtClient(mocks.NewDatabaseClient(t))

	// then
	require.NoError(t, err)
	assert.NotNil(t, client.Client)
	assert.NotNil(t, client.QueryTemplater)
}
