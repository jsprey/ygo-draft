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

//go:embed test_templates/testQuerySelectAllSets.sql
var testQuerySelectAllSets string

func TestNewSqlQueryTemplater_QuerySelectAllSets(t *testing.T) {
	t.Run("correctly create query", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectAllSets()

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllSets, myString)
	})
}

//go:embed test_templates/testQuerySelectAllCardsWithFilter_NoFilter.sql
var testQuerySelectAllCardsWithFilter_NoFilter string

//go:embed test_templates/testQuerySelectAllCardsWithFilter_WithTypeFilter.sql
var testQuerySelectAllCardsWithFilter_WithTypeFilter string

//go:embed test_templates/testQuerySelectAllCardsWithFilter_WithSetFilter.sql
var testQuerySelectAllCardsWithFilter_WithSetFilter string

//go:embed test_templates/testQuerySelectAllCardsWithFilter_WithAllFilter.sql
var testQuerySelectAllCardsWithFilter_WithAllFilter string

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

	t.Run("correctly create query with types", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		filter := model.CardFilter{
			Types: []string{"Normal Monster", "Effect Monster"},
			Sets:  nil,
		}

		// when
		myString, err := templater.SelectAllCardsWithFilter(filter)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_WithTypeFilter, myString)
	})

	t.Run("correctly create query with sets", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		filter := model.CardFilter{
			Types: []string{},
			Sets:  []string{"Force of the Breaker", "Strike of Neos"},
		}

		// when
		myString, err := templater.SelectAllCardsWithFilter(filter)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_WithSetFilter, myString)
	})

	t.Run("correctly create query with set and type filter", func(t *testing.T) {
		// given
		templater, err := query.NewSqlQueryTemplater()
		require.NoError(t, err)
		filter := model.CardFilter{
			Types: []string{"Normal Monster", "Effect Monster"},
			Sets:  []string{"Force of the Breaker", "Strike of Neos"},
		}

		// when
		myString, err := templater.SelectAllCardsWithFilter(filter)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectAllCardsWithFilter_WithAllFilter, myString)
	})
}
