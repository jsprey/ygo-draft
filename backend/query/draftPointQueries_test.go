package query

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
)

//go:embed test_templates/drafts/points/testQuerySelectPoints.sql
var testQuerySelectPoints string

func Test_sqlQueryTemplater_SelectPoints(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		// given
		templater, err := NewSqlQueryTemplater()
		require.NoError(t, err)

		product := model.DraftStoreProduct{
			ID:          "test_product",
			DisplayName: "Draft One Card",
			Description: "You can draft another card for the next round.",
			Rarity:      model.ProductRarityCommon,
			Costs:       123,
		}
		product.SetTags(model.ProductTagPositive, model.ProductTagGainCard)

		// when
		myString, err := templater.SelectPoints(4, 3)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectPoints, myString)
	})
}

//go:embed test_templates/drafts/points/testQueryUpsertPoints.sql
var testQueryUpsertPoints string

func Test_sqlQueryTemplater_UpsertPoints(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		// given
		templater, err := NewSqlQueryTemplater()
		require.NoError(t, err)

		product := model.DraftStoreProduct{
			ID:          "test_product",
			DisplayName: "Draft One Card",
			Description: "You can draft another card for the next round.",
			Rarity:      model.ProductRarityCommon,
			Costs:       123,
		}
		product.SetTags(model.ProductTagPositive, model.ProductTagGainCard)

		// when
		myString, err := templater.UpsertPoints(4, 5, 55)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUpsertPoints, myString)
	})
}
