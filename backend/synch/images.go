package synch

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"ygodraft/backend/model"
)

const ImagesDirectoryName = "imageStore"

func (yds *YgoDataSyncher) SynchImageOfCard(card *model.Card) error {
	if len(card.Images) == 0 {
		logrus.Warnf("Card [%d][%s]", card.ID, card.Name)
		return nil
	}

	err := fetchAndSaveImage(card.ID, "small.png", card.Images[0].ImageURLSmall)
	if err != nil {
		return fmt.Errorf("failed to fetch and save big image for card [%d]: %w", card.ID, err)
	}

	err = fetchAndSaveImage(card.ID, "big.png", card.Images[0].ImageURL)
	if err != nil {
		return fmt.Errorf("failed to fetch and save big image for card [%d]: %w", card.ID, err)
	}

	return nil
}

func fetchAndSaveImage(id int, fileName string, imageUrl string) error {
	err := ensurePathExists(getFolderPath(id))
	if err != nil {
		return fmt.Errorf("failed to ensure folder for [%s]: %w", getFolderPath(id), err)
	}

	imagePath := fmt.Sprintf("%s%c%s", getFolderPath(id), os.PathSeparator, fileName)
	exists, err := isFileExisting(imagePath)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	imageBytes, err := getImage(imageUrl)
	if err != nil {
		return err
	}

	err = saveImage(imagePath, imageBytes)
	if err != nil {
		return err
	}

	return nil
}

func getImage(url string) ([]byte, error) {
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get [%s]: %w", url, err)
	}

	return ioutil.ReadAll(response.Body)
}

func saveImage(filePath string, image []byte) error {
	err := ioutil.WriteFile(filePath, image, 0744)
	if err != nil {
		return fmt.Errorf("failed to save image [%s]: %w", filePath, err)
	}

	return nil
}

func isFileExisting(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("failed to stat [%s]: %w", path, err)
	}

	return !os.IsNotExist(err), nil
}

func getFolderPath(cardID int) string {
	return fmt.Sprintf(".%c%s%c%d", os.PathSeparator, ImagesDirectoryName, os.PathSeparator, cardID)
}

func ensurePathExists(folder string) error {
	fileStat, err := os.Stat(folder)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	if os.IsNotExist(err) {
		logrus.Tracef("Creating directory [%s]", folder)

		err := os.MkdirAll(folder, 0744)
		if err != nil {
			return fmt.Errorf("failed to create directory [%s]: %w", folder, err)
		}
	} else {
		if !fileStat.IsDir() {
			return fmt.Errorf("cannot store images in directory [%s] as it already exists and is a file", fileStat.Sys())
		}
	}

	return nil
}
