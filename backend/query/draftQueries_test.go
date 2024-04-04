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

func Test_newSqlQueryTemplater_QueryInsertDraft(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.InsertDraft(4, 5, 6, model.DraftStatusRunning)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertDraft, myString)
	})
}

//go:embed test_templates/drafts/testQueryUpdateDraft.sql
var testQueryUpdateDraft string

func Test_newSqlQueryTemplater_QueryUpdateDraft(t *testing.T) {
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
