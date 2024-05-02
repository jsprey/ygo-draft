package model

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

// ProductRarity determines the rarity of the draft refinement product.
type ProductRarity string

const (
	ProductRarityCommon    ProductRarity = "common"
	ProductRarityRare      ProductRarity = "rare"
	ProductRarityEpic      ProductRarity = "epic"
	ProductRarityLegendary ProductRarity = "legendary"
)

// DraftStoreProduct represents an entry in the store.
type DraftStoreProduct struct {
	ID          string        `json:"id"`
	DisplayName string        `json:"display_name"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Rarity      ProductRarity `json:"rarity"`
	Tags        string        `json:"tags"`
	Costs       int           `json:"costs"`
}

// DraftStoreCatalog contains all store products.
type DraftStoreCatalog struct {
	Products []*DraftStoreProduct `json:"products"`
}

const tagSeparator = ","
const (
	ProductTagGainCard string = "gain_card"
	ProductTagLoseCard string = "lose_card"
	ProductTagPositive string = "positive_effect"
	ProductTagNegative string = "negative_effect"
)

// GetTags returns the tags of the product as a slice
func (d *DraftStoreProduct) GetTags() []string {
	return strings.Split(d.Tags, tagSeparator)
}

// SetTags sets the tags for the given element
func (d *DraftStoreProduct) SetTags(tags ...string) {
	d.Tags = strings.Join(tags, tagSeparator)
}

// HasProduct checks if the catalog has a product with a given id.
func (d *DraftStoreCatalog) HasProduct(productId string) bool {
	for _, product := range d.Products {
		if product.ID == productId {
			return true
		}
	}

	return false
}

// GetProduct retrieves a product from the catalog given by the id.
func (d *DraftStoreCatalog) GetProduct(productID string) (*DraftStoreProduct, error) {
	for _, product := range d.Products {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, fmt.Errorf("there is no product with the id %s in the catalog", productID)
}

// RemoveProduct removes a product from the catalog.
func (d *DraftStoreCatalog) RemoveProduct(product *DraftStoreProduct) {
	if !d.HasProduct(product.ID) {
		return
	}

	d.Products = slices.DeleteFunc(d.Products, func(p *DraftStoreProduct) bool {
		return p == product // delete the odd numbers
	})
}

// AddProduct adds a new product to the catalog.
func (d *DraftStoreCatalog) AddProduct(product *DraftStoreProduct) {
	if d.HasProduct(product.ID) {
		return
	}

	d.Products = append(d.Products, product)
}

// Sort sorts the products of the catalog.
func (d *DraftStoreCatalog) Sort() {
	slices.SortFunc(d.Products, func(a, b *DraftStoreProduct) int {
		return cmp.Compare(a.ID, b.ID)
	})
}

// DraftStoreClient provides the functionality to manage the store in drafts.
type DraftStoreClient interface {
	// GetProducts retrieves all available products from the database.
	GetProducts() (*DraftStoreCatalog, error)
	// BuyProduct buys a product for a user in a certain draft.
	BuyProduct(draftID int, userID int, productID int) error
	// SyncCatalog updates the catalog of products to the newest state.
	SyncCatalog() error
}

// DraftStoreQueryGenerator is responsible to generate queries related to the draft store.
type DraftStoreQueryGenerator interface {
	// UpsertStoreProduct creates an upsert query to create a new store product.
	UpsertStoreProduct(product DraftStoreProduct) (string, error)
	// DeleteStoreProduct creates a delete query to delete a store product.
	DeleteStoreProduct(product DraftStoreProduct) (string, error)
	// SelectStoreProduct creates a select query to retrieve all available products.
	SelectStoreProduct() (string, error)
}
