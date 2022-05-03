package query_test

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

//go:embed test_templates/QueryInsertCard.sql
var testQueryInsertCard string

func TestQueryInsertCard(t *testing.T) {
	t.Run("correctly escape my things", func(t *testing.T) {
		// given
		myCard := &model.Card{
			Name:      "My super ' duper '' Name '''",
			Type:      "My super % duper %% Name %%%",
			Desc:      "My super ¹²³¼½¬{[]} duper ¹²³¼½¬{[]} Name ¹²³¼½¬{[]}",
			Race:      "My super \\ duper \\\\ Name \\\\\\",
			Attribute: "My super \" duper \"\" Name \"\"\"",
		}

		// when
		myString, err := query.QueryInsertCard(myCard)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryInsertCard, myString)
	})
}
