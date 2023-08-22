package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/query"
)

//go:embed test_templates/users/testQuerySelectUserByEmail.sql
var testQueryUsersSelectUserByEmail string

func TestNewSqlQueryTemplater_QuerySelectUserByEmail(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		myEmail := "test@test.de"

		// when
		myString, err := templater.SelectUserByEmail(myEmail)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUsersSelectUserByEmail, myString)
	})
}

//go:embed test_templates/users/testQueryInsertUserAdmin.sql
var testQueryUsersInsertUser string

func TestNewSqlQueryTemplater_QueryInsertUser(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// given
		myEmail := "test@test.de"
		myPasswordHash := "MEGAHASH"
		myDisplayName := "MyDisplayName"

		// when
		myString, err := templater.InsertUser(myEmail, myPasswordHash, myDisplayName, true)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUsersInsertUser, myString)
	})
}
