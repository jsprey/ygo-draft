package query

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/model"
)

//go:embed test_templates/drafts/store/testQuerySelectStoreProduct.sql
var testQuerySelectStoreProduct string

func Test_sqlQueryTemplater_SelectStoreProduct(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		// given
		templater, err := NewSqlQueryTemplater()
		require.NoError(t, err)

		// when
		myString, err := templater.SelectStoreProduct()

		// then
		require.NoError(t, err)
		assert.Equal(t, testQuerySelectStoreProduct, myString)
	})
}

//go:embed test_templates/drafts/store/testQueryUpsertStoreProduct.sql
var testQueryUpsertStoreProduct string

func Test_sqlQueryTemplater_UpsertStoreProduct(t *testing.T) {
	t.Run("create query", func(t *testing.T) {
		// given
		templater, err := NewSqlQueryTemplater()
		require.NoError(t, err)

		product := model.DraftStoreProduct{
			ID:          "draw_one_more",
			DisplayName: "Draft One Card",
			Description: "You can draft another card for the next round.",
			Category:    "Extra Draws",
			Rarity:      model.ProductRarityCommon,
			Costs:       123,
		}
		product.SetTags(model.ProductTagPositive, model.ProductTagGainCard)

		// when
		myString, err := templater.UpsertStoreProduct(product)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryUpsertStoreProduct, myString)
	})
}

//go:embed test_templates/drafts/store/testQueryDeleteStoreProduct.sql
var testQueryDeleteStoreProduct string

func Test_sqlQueryTemplater_DeleteStoreProduct(t *testing.T) {
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
		myString, err := templater.DeleteStoreProduct(product)

		// then
		require.NoError(t, err)
		assert.Equal(t, testQueryDeleteStoreProduct, myString)
	})
}
