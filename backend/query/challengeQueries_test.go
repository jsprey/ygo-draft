package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/challenges/testQuerySelectOutgoingChallenges.sql
var testQuerySelectOutgoingChallenges string

//go:embed test_templates/challenges/testQuerySelectOutgoingChallenges_AllStatus.sql
var testQuerySelectOutgoingChallengesAllStatus string

func Test_sqlQueryTemplater_SelectOutgoingChallenges(t *testing.T) {
	t.Run("create query with status", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectOutgoingChallenges(0, model.StatusPending)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectOutgoingChallenges, myString)
	})
	t.Run("create query with no required status", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectOutgoingChallenges(0, model.StatusAll)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectOutgoingChallengesAllStatus, myString)
	})
}

//go:embed test_templates/challenges/testQuerySelectReceivedChallenges.sql
var testQuerySelectReceivedChallenges string

//go:embed test_templates/challenges/testQuerySelectReceivedChallenges_AllStatus.sql
var testQuerySelectReceivedChallengesAllStatus string

func Test_sqlQueryTemplater_SelectReceivedChallenges(t *testing.T) {
	t.Run("create query with status", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectReceivedChallenges(3, model.StatusPending)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectReceivedChallenges, myString)
	})
	t.Run("create query with no required status", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectReceivedChallenges(3, model.StatusAll)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectReceivedChallengesAllStatus, myString)
	})
}

//go:embed test_templates/challenges/testQuerySelectChallenge.sql
var testQuerySelectChallenge string

func Test_sqlQueryTemplater_SelectChallenge(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectChallenge(3)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectChallenge, myString)
	})
}

//go:embed test_templates/challenges/testQueryUpdateChallenge.sql
var testQueryUpdateChallenge string

func Test_sqlQueryTemplater_UpdateChallenge(t *testing.T) {
	t.Run("fail to create query as invalid status was provided", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		_, err = templater.UpdateChallenge(4, model.StatusAll)

		// then
		require.Error(t, err)
		require.ErrorContains(t, err, "invalid status")
	})
	t.Run("create query successful", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.UpdateChallenge(4, model.StatusPending)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUpdateChallenge, myString)
	})
}

//go:embed test_templates/challenges/testQueryInsertChallenge.sql
var testQueryInsertChallenge string

func Test_sqlQueryTemplater_InsertChallenge(t *testing.T) {
	set1 := model.CardSet{SetName: "Set 1", SetCode: "SS-01", SetRarity: "rare", SetRarityCode: "R"}
	set2 := model.CardSet{SetName: "Set S", SetCode: "TS-01", SetRarity: "common", SetRarityCode: "C"}
	settings := model.DraftSettings{
		MainDeckDraws:  40,
		ExtraDeckDraws: 3,
		Mode:           model.DraftModeBestOf,
		ModeValue:      5,
		Sets:           []model.CardSet{set1, set2},
	}

	t.Run("create a new challenge", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.InsertChallenge(4, 6, settings)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertChallenge, myString)
	})
}
