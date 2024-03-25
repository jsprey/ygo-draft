package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/users/testQuerySelectUserByID.sql
var testQueryUsersSelectUserByID string

func Test_newSqlQueryTemplater_QuerySelectUserByID(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectUserByID(0)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUsersSelectUserByID, myString)
	})
}

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
		user := model.User{
			Email:        "test@test.de",
			PasswordHash: "MEGAHASH",
			DisplayName:  "MyDisplayName",
			IsAdmin:      true,
		}

		// when
		myString, err := templater.InsertUser(user)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUsersInsertUser, myString)
	})
}

//go:embed test_templates/users/testQueryDeleteUser.sql
var testQueryUsersDeleteUser string

func TestNewSqlQueryTemplater_QueryDeleteUser(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// given
		myEmail := "test@test.de"

		// when
		myString, err := templater.DeleteUser(myEmail)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUsersDeleteUser, myString)
	})
}

//go:embed test_templates/users/testQuerySelectUsers.sql
var testQuerySelectUsers string

func Test_newSqlQueryTemplater_SelectUsers(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectUsers(1, 20)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectUsers, myString)
	})
}
