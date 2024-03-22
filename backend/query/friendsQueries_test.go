package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/users/friends/testQueryGetFriends.sql
var testQueryGetFriends string

func Test_sqlQueryTemplater_QueryGetFriends(t *testing.T) {
	t.Run("wrongly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.GetFriends(0)

		// then
		require.NoError(t, err)
		assert.NotEqual(t, testQueryGetFriends, myString)
	})
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.GetFriends(5)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryGetFriends, myString)
	})
}

//go:embed test_templates/users/friends/testQueryGetFriendRequests.sql
var testQueryGetFriendRequests string

func Test_sqlQueryTemplater_GetFriendRequests(t *testing.T) {
	t.Run("wrongly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.GetFriendRequests(9)

		// then
		require.NoError(t, err)
		assert.NotEqual(t, testQueryGetFriendRequests, myString)
	})
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.GetFriendRequests(5)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryGetFriendRequests, myString)
	})
}

//go:embed test_templates/users/friends/testQuerySetFriendRelation.sql
var testQuerySetFriendRelation string

func Test_sqlQueryTemplater_SetFriendRelation(t *testing.T) {
	t.Run("wrongly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SetFriendRelation(5, 6, model.FriendStatusInvited)

		// then
		require.NoError(t, err)
		assert.NotEqual(t, testQuerySetFriendRelation, myString)
	})
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SetFriendRelation(5, 6, model.FriendStatusFriends)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySetFriendRelation, myString)
	})
}
