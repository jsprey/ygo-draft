package setup

import (
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"ygodraft/backend/client/draft"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
)

//go:embed queries/createTables.sql
var createTablesQuery string

// DatabaseSetup is responsible to setup the database including the creation of the database and the data tables.
type DatabaseSetup struct {
	Client            model.DatabaseClient
	DraftClient       model.DraftClient
	DraftPointsClient model.DraftPointsClient
	DraftStoreClient  model.DraftStoreClient
	UsermgtClient     model.UsermgtClient
}

func NewDatabaseSetup(client model.DatabaseClient, usermgtClient model.UsermgtClient) (*DatabaseSetup, error) {
	draftClient, err := draft.NewDraftClient(client)
	if err != nil {
		return nil, fmt.Errorf("failed to draft client: %w", err)
	}

	draftPointClient, err := draft.NewDraftPointsClient(client, draftClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create draft points client: %w", err)
	}

	draftStoreClient, err := draft.NewDraftStoreClient(client, draftClient, draftPointClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create draft store client: %w", err)
	}

	return &DatabaseSetup{Client: client,
		DraftClient:       draftClient,
		DraftPointsClient: draftPointClient,
		DraftStoreClient:  draftStoreClient,
		UsermgtClient:     usermgtClient,
	}, nil
}

func (ds *DatabaseSetup) Setup() error {
	logrus.Debug("Setup -> Database -> Creating tables")
	_, err := ds.Client.Exec(createTablesQuery)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	err = ds.setupUsermgt()
	if err != nil {
		return fmt.Errorf("failed to setup usermgt database stuff: %w", err)
	}

	err = ds.setupProductStore()
	if err != nil {
		return fmt.Errorf("failed to setup product store: %w", err)
	}

	return nil
}

func (ds *DatabaseSetup) setupProductStore() error {
	return ds.DraftStoreClient.SyncCatalog()
}

func (ds *DatabaseSetup) setupUsermgt() error {
	logrus.Debugf("Creating default admin user...")
	err := createUserIfNotExist(ds.UsermgtClient, config.AdminUserEmail, "Admin", true, "adminadmin")
	if err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	return nil
}

func createUserIfNotExist(client model.UsermgtClient, email string, displayName string, isAdmin bool, clearTextPassword string) error {
	_, err := client.GetUser(email)
	if err != nil && !model.IsErrorUserDoesNotExist(err) {
		return fmt.Errorf("failed to get user: %w", err)
	} else if model.IsErrorUserDoesNotExist(err) {
		logrus.Warningf("Creating user [%s]", email)

		// hash password from config bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clearTextPassword), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to generate password hash: %w", err)
		}

		// create admin user
		user := model.User{
			Email:        email,
			PasswordHash: string(hashedPassword),
			DisplayName:  displayName,
			IsAdmin:      isAdmin,
		}

		err = client.CreateUser(user)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	}

	return nil
}
