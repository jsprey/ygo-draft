package query

import (
	_ "embed"
	"ygodraft/backend/model"
)

func (sqt *sqlQueryTemplater) AddDraftStoreQueries(templateMap *map[string]string) {
	(*templateMap)["SelectStoreProduct"] = templateContentSelectStoreProduct
	(*templateMap)["UpsertStoreProduct"] = templateContentUpsertStoreProduct
	(*templateMap)["DeleteStoreProduct"] = templateContentDeleteStoreProduct
}

//go:embed templates/drafts/store/QuerySelectStoreProduct.sql
var templateContentSelectStoreProduct string

func (sqt *sqlQueryTemplater) SelectStoreProduct() (string, error) {
	return sqt.Template("SelectStoreProduct", "")
}

//go:embed templates/drafts/store/QueryUpsertStoreProduct.sql
var templateContentUpsertStoreProduct string

func (sqt *sqlQueryTemplater) UpsertStoreProduct(product model.DraftStoreProduct) (string, error) {
	templateObject := struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Rarity      string `json:"rarity"`
		Tags        string `json:"tags"`
		Costs       int    `json:"costs"`
	}{
		ID:          escape(product.ID),
		DisplayName: escape(product.DisplayName),
		Description: escape(product.Description),
		Category:    escape(product.Category),
		Rarity:      escape(string(product.Rarity)),
		Tags:        escape(product.Tags),
		Costs:       product.Costs,
	}

	return sqt.Template("UpsertStoreProduct", &templateObject)
}

//go:embed templates/drafts/store/QueryDeleteStoreProduct.sql
var templateContentDeleteStoreProduct string

func (sqt *sqlQueryTemplater) DeleteStoreProduct(product model.DraftStoreProduct) (string, error) {
	templateObject := struct {
		ID string `json:"id"`
	}{
		ID: escape(product.ID),
	}

	return sqt.Template("DeleteStoreProduct", &templateObject)
}
