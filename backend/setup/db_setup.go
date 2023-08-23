package setup

import (
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"ygodraft/backend/client/usermgt"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
)

//go:embed queries/cards.sql
var createTableCards string

//go:embed queries/card_sets.sql
var createTableCardSets string

//go:embed queries/users.sql
var createTableUsers string

// DatabaseSetup is responsible to setup the database including the creation of the database and the data tables.
type DatabaseSetup struct {
	Client      model.DatabaseClient
	authContext config.AuthContext
}

func NewDatabaseSetup(client model.DatabaseClient, context config.AuthContext) *DatabaseSetup {
	return &DatabaseSetup{Client: client, authContext: context}
}

func (ds *DatabaseSetup) Setup() error {
	logrus.Debug("Setup -> Database -> Creating `cards` table")
	_, err := ds.Client.Exec(createTableCards)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	logrus.Debug("Setup -> Database -> Creating `card_sets` table")
	_, err = ds.Client.Exec(createTableCardSets)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	logrus.Debug("Setup -> Database -> Creating `users` table")
	_, err = ds.Client.Exec(createTableUsers)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	err = ds.setupUsermgt(err)
	if err != nil {
		return fmt.Errorf("failed to setup usermgt database stuff: %w", err)
	}

	return nil
}

func (ds *DatabaseSetup) setupUsermgt(err error) error {
	// create default admin user if not exists
	client, err := usermgt.NewUsermgtClient(ds.Client)
	if err != nil {
		return fmt.Errorf("failed to create new usermgt client: %w", err)
	}

	_, err = client.GetUser(config.AdminUserEmail)
	if err != nil && !model.IsErrorUserDoesNotExist(err) {
		return fmt.Errorf("failed to get user: %w", err)
	} else if model.IsErrorUserDoesNotExist(err) {
		logrus.Debugf("Default Admin does not exist -> Creating default admin user...")
		// hash password from config bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ds.authContext.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to generate password hash: %w", err)
		}

		// create admin user
		adminUser := model.User{
			Email:        config.AdminUserEmail,
			PasswordHash: string(hashedPassword),
			DisplayName:  "Admin",
			IsAdmin:      true,
		}

		err = client.CreateUser(adminUser)
		if err != nil {
			return fmt.Errorf("failed to create admin user: %w", err)
		}
	}

	return nil
}
