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
