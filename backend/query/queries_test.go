package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/testQueryInsertCard.sql
var testQueryInsertCard string

func TestNewSqlQueryTemplater_QueryInsertCard(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// given
		myCard := &model.Card{
			Name:      "My super ' duper '' Name '''",
			Type:      "My super % duper %% Name %%%",
			Desc:      "My super ¹²³¼½¬{[]} duper ¹²³¼½¬{[]} Name ¹²³¼½¬{[]}",
			Race:      "My super \\ duper \\\\ Name \\\\\\",
			Attribute: "My super \" duper \"\" Name \"\"\"",
		}

		// when
		myString, err := templater.InsertCard(myCard)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertCard, myString)
	})
}

//go:embed test_templates/testQuerySelectCardByID.sql
var testQuerySelectCardByID string

func TestNewSqlQueryTemplater_QuerySelectCardByID(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectCardByID(1231234)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectCardByID, myString)
	})
}

//go:embed test_templates/testQuerySelectAllCards.sql
var testQuerySelectAllCards string

func TestNewSqlQueryTemplater_QuerySelectAllCards(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectAllCards()

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCards, myString)
	})
}

//go:embed test_templates/testQuerySelectAllCardsWithFilter_NoFilter.sql
var testQuerySelectAllCardsWithFilter_NoFilter string

//go:embed test_templates/testQuerySelectAllCardsWithFilter_WithFilter.sql
var testQuerySelectAllCardsWithFilter_WithFilter string

func TestNewSqlQueryTemplater_QuerySelectAllCardsWithFilter(t *testing.T) {
	t.Run("correctly create query with empty filter", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectAllCardsWithFilter(model.CardFilter{})

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_NoFilter, myString)
	})

	t.Run("correctly create query with filters", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		filter := model.CardFilter{
			Types: []string{"HAHA", "Second Type", "third type"},
			Sets:  nil,
		}

		// when
		myString, err := templater.SelectAllCardsWithFilter(filter)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_WithFilter, myString)
	})

	t.Run("correctly create query with filters and sets", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		filter := model.CardFilter{
			Types: []string{"HAHA", "Second Type", "third type"},
			Sets:  []string{}, //todo add sets
		}

		// when
		myString, err := templater.SelectAllCardsWithFilter(filter)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_WithFilter, myString)
	})
}
