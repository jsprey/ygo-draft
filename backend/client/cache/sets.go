package cache

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
)

func (yc *YgoCache) SaveSets(sets []model.CardSet) error {
	numberOfSets := len(sets)
	logrus.Infof("Cache -> Starting the synchronization of %d sets...", numberOfSets)

	for _, set := range sets {
		err := yc.saveSet(set)
		if err != nil {
			return fmt.Errorf("failed to save set [%s]: %w", set.SetCode, err)
		}
	}

	return nil
}

func (yc *YgoCache) GetAllSets() (*[]*model.CardSet, error) {
	logrus.Debug("Cache -> Retrieve all sets")

	sqlQuery, err := yc.QueryTemplater.SelectAllSets()
	if err != nil {
		return nil, err
	}

	var sets []*model.CardSet
	err = yc.Client.Select(sqlQuery, &sets)
	if err != nil {
		return nil, fmt.Errorf("failed to scan struct: %w", err)
	}

	return &sets, nil
}

func (yc *YgoCache) saveSet(set model.CardSet) error {
	logrus.Debugf("Cache -> Save set with id %s", set.SetCode)
	sqlQuery, err := yc.QueryTemplater.InsertSet(set)
	if err != nil {
		return err
	}

	_, err = yc.Client.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [%s]: %w", sqlQuery, err)
	}

	return nil
}
