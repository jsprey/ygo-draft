package draft

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type draftStoreClient struct {
	Client           model.DatabaseClient
	DraftClient      model.DraftClient
	DraftPointClient model.DraftPointsClient
	QueryTemplater   model.DraftStoreQueryGenerator
}

// NewDraftStoreClient creates a new instance of the draft store client.
func NewDraftStoreClient(dbClient model.DatabaseClient, draftClient model.DraftClient, draftPointsClient model.DraftPointsClient) (*draftStoreClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &draftStoreClient{
		Client:           dbClient,
		DraftClient:      draftClient,
		DraftPointClient: draftPointsClient,
		QueryTemplater:   queryTemplater,
	}, nil
}

func (d draftStoreClient) GetProducts() (*model.DraftStoreCatalog, error) {
	selectStoreProductQuery, err := d.QueryTemplater.SelectStoreProduct()
	if err != nil {
		return nil, fmt.Errorf("failed to template query [SelectStoreProduct]: %w", err)
	}

	var selectedRows []*model.DraftStoreProduct
	err = d.Client.Select(selectStoreProductQuery, &selectedRows)
	if err != nil {
		return nil, fmt.Errorf("failed to select query [SelectStoreProduct]: %w", err)
	}

	if selectedRows == nil || len(selectedRows) == 0 {
		return &model.DraftStoreCatalog{Products: []*model.DraftStoreProduct{}}, nil
	}

	return &model.DraftStoreCatalog{Products: selectedRows}, nil
}

func (d draftStoreClient) BuyProduct(draftID int, userID int, productID int) error {
	//TODO implement me
	panic("implement me")
}

func (d draftStoreClient) SyncCatalog() error {
	fileCatalog, err := d.getProductCatalogFromFile()
	if err != nil {
		return fmt.Errorf("failed to get product catalog from file: %w", err)
	}

	dbCatalog, err := d.GetProducts()
	if err != nil {
		return fmt.Errorf("failed to get product catalog from database: %w", err)
	}

	err = d.deleteProducts(fileCatalog, dbCatalog)
	if err != nil {
		return fmt.Errorf("failed to delete no longer available products: %w", err)
	}

	err = d.upsertProducts(fileCatalog.Products)
	if err != nil {
		return fmt.Errorf("failed to update products: %w", err)
	}

	return nil
}

func (d draftStoreClient) deleteProducts(fileCatalog *model.DraftStoreCatalog, dbCatalog *model.DraftStoreCatalog) error {
	var multiErr error
	toBeDeleted := &model.DraftStoreCatalog{Products: []*model.DraftStoreProduct{}}
	for _, actualProduct := range dbCatalog.Products {
		if !fileCatalog.HasProduct(actualProduct.ID) {
			product, err := dbCatalog.GetProduct(actualProduct.ID)
			if err != nil {
				return fmt.Errorf("failed to get product: %w", err)
			}

			toBeDeleted.AddProduct(product)
		}
	}
	if multiErr != nil {
		return multiErr
	}

	for _, product := range toBeDeleted.Products {
		deleteProductQuery, err := d.QueryTemplater.DeleteStoreProduct(*product)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("failed to template query [DeleteStoreProduct]: %w", err))
			continue
		}

		_, err = d.Client.Exec(deleteProductQuery)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("failed to exec query [DeleteStoreProduct]: %w", err))
			continue
		}
	}

	return multiErr
}

func (d draftStoreClient) upsertProducts(products []*model.DraftStoreProduct) error {
	var multiErr error
	for _, product := range products {
		upsertProductQuery, err := d.QueryTemplater.UpsertStoreProduct(*product)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("failed to template query [DeleteStoreProduct]: %w", err))
			continue
		}

		_, err = d.Client.Exec(upsertProductQuery)
		if err != nil {
			multiErr = multierror.Append(multiErr, fmt.Errorf("failed to exec query [DeleteStoreProduct]: %w", err))
			continue
		}
	}

	return multiErr
}

func (d draftStoreClient) getProductCatalogFromFile() (*model.DraftStoreCatalog, error) {
	catalogFolderPath := path.Dir(config.ProductCatalogFilePath)
	_, err := os.Stat(catalogFolderPath)
	if os.IsNotExist(err) {
		logrus.Debugf("No data folder exists -> create")
		err := os.MkdirAll(catalogFolderPath, os.ModeDir|os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to mkdirall: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}

	// check if catalog exists
	fileContent, err := os.ReadFile(config.ProductCatalogFilePath)
	if os.IsNotExist(err) {
		logrus.Debugf("No product catalog file exists -> create")
		_, err = os.Create(config.ProductCatalogFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create product catalog file: %w", err)
		}

		// create simple product example
		product1 := model.DraftStoreProduct{
			ID: "new", DisplayName: "New Product", Description: "This is new", Rarity: model.ProductRarityCommon, Costs: 500,
		}
		product1.SetTags(model.ProductTagLoseCard)
		product2 := model.DraftStoreProduct{
			ID: "new2", DisplayName: "New Product2", Description: "This is new 2", Rarity: model.ProductRarityCommon, Costs: 500,
		}
		product2.SetTags(model.ProductTagGainCard)
		catalog := &model.DraftStoreCatalog{
			Products: []*model.DraftStoreProduct{&product1, &product2},
		}

		catalogXmlData, err := yaml.Marshal(catalog)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal product catalog: %w", err)
		}

		err = os.WriteFile(config.ProductCatalogFilePath, catalogXmlData, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to write simple product catalog: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}

	toBeUpserted := &model.DraftStoreCatalog{Products: []*model.DraftStoreProduct{}}
	err = yaml.Unmarshal(fileContent, toBeUpserted)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product catalog: %w", err)
	}

	return toBeUpserted, nil
}
