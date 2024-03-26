package setup

import (
	_ "embed"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"ygodraft/backend/config"
	"ygodraft/backend/model"
)

//go:embed queries/cards.sql
var createTableCards string

//go:embed queries/card_sets.sql
var createTableCardSets string

//go:embed queries/users.sql
var createTableUsers string

//go:embed queries/friends.sql
var createTableFriends string

// DatabaseSetup is responsible to setup the database including the creation of the database and the data tables.
type DatabaseSetup struct {
	Client        model.DatabaseClient
	UsermgtClient model.UsermgtClient
}

func NewDatabaseSetup(client model.DatabaseClient, usermgtClient model.UsermgtClient) *DatabaseSetup {
	return &DatabaseSetup{Client: client, UsermgtClient: usermgtClient}
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

	logrus.Debug("Setup -> Database -> Creating `friends` table")
	_, err = ds.Client.Exec(createTableFriends)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	err = ds.setupUsermgt()
	if err != nil {
		return fmt.Errorf("failed to setup usermgt database stuff: %w", err)
	}

	return nil
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
