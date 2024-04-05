package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/drafts/testQueryInsertDraft.sql
var testQueryInsertDraft string

func Test_newSqlQueryTemplater_InsertDraft(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  4,
			MainDeckSize:   5,
			ExtraDeckDraws: 6,
			ExtraDeckSize:  7,
			Mode:           "rounds",
			ModeValue:      10,
			Sets: []model.CardSet{
				{SetName: "Test-Set", SetCode: "CCSR-3", SetRarity: "RR", SetRarityCode: "CC"},
			},
		}

		// when
		myString, err := templater.InsertDraft(4, 5, settings)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertDraft, myString)
	})
}

//go:embed test_templates/drafts/testQueryUpdateDraft.sql
var testQueryUpdateDraft string

func Test_newSqlQueryTemplater_UpdateDraft(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.UpdateDraft(4, model.DraftStatusRunning)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUpdateDraft, myString)
	})
}

//go:embed test_templates/drafts/rounds/testQueryInsertDraftRound.sql
var testQueryInsertDraftRound string

func Test_newSqlQueryTemplater_InsertDraftRound(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.InsertDraftRound(4, 5, model.DraftRoundStatusPreparation)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertDraftRound, myString)
	})
}

//go:embed test_templates/drafts/rounds/testQueryUpdateDraftRound.sql
var testQueryUpdateDraftRound string

func Test_newSqlQueryTemplater_UpdateDraftRound(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.UpdateDraftRound(4, 5, model.DraftRoundStatusFinished)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUpdateDraftRound, myString)
	})
}

//go:embed test_templates/drafts/deck/testQueryInsertDraftRoundDeck.sql
var testQueryInsertDraftRoundDeck string

func Test_newSqlQueryTemplater_InsertDraftRoundDeck(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		deck := []string{
			"123", "132", "141",
		}

		// when
		myString, err := templater.InsertDraftRoundDeck(4, 5, deck)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertDraftRoundDeck, myString)
	})
}

//go:embed test_templates/drafts/testQuerySelectDraft.sql
var testQuerySelectDraft string

func Test_sqlQueryTemplater_SelectDrafts(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectDraft(4)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectDraft, myString)
	})
}

//go:embed test_templates/drafts/testQuerySelectDraftsWithStatus.sql
var testQuerySelectDraftsWithStatus string

func Test_sqlQueryTemplater_SelectDraftsWithStatus(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectDraftsWithStatus(4, model.DraftStatusRunning)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectDraftsWithStatus, myString)
	})
}

//go:embed test_templates/drafts/testQuerySelectDraftsWithUsersAndStatus.sql
var testQuerySelectDraftsWithUsersAndStatus string

func Test_sqlQueryTemplater_SelectDraftsWithUsersAndStatus(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectDraftsWithUsersAndStatus(1, 2, model.DraftStatusRunning)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectDraftsWithUsersAndStatus, myString)
	})
}
