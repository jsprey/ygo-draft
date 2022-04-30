package ygoprodeck

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"ygodraft/backend/model"
)

const endpointCardInfo = "https://db.ygoprodeck.com/api/v7"

// YgoProDeckClient is responsible to extract all necessary card information from the ygoprodeck website.
type YgoProDeckClient struct {
	BaseUrl string        `json:"base_url,omitempty"`
	Client  *RLHTTPClient `json:"Client" json:"client,omitempty"`
}

// NewYgoProDeckClient creates a new instance of the YgoProDeckClient.
func NewYgoProDeckClient() *YgoProDeckClient {
	rateLimitedClient := NewDefaultRateLimitedClient()

	return &YgoProDeckClient{
		BaseUrl: endpointCardInfo,
		Client:  rateLimitedClient,
	}
}

// GetAllCards retrieves all cards from the ygo pro deck api.
func (ypdc YgoProDeckClient) GetAllCards() (*[]*model.Card, error) {
	targetUrl := fmt.Sprintf("%s/cardinfo.php", ypdc.BaseUrl)
	logrus.Debugf("YgoProDeckClient -> GetAllCards -> Requesting [%s]", targetUrl)

	resp, err := ypdc.Client.Client.Get(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get [%s]: %w", ypdc.BaseUrl, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read body: %w", err)
	}

	var cards struct {
		Data []*model.Card `json:"data"`
	}
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return &cards.Data, nil
}

// GetCard retrieves a card with the given id from the ygo pro deck api.
func (ypdc YgoProDeckClient) GetCard(id int) (*model.Card, error) {
	targetUrl := fmt.Sprintf("%s/cardinfo.php?id=%d", ypdc.BaseUrl, id)
	logrus.Debugf("YgoProDeckClient -> GetCard [%d] -> Requesting [%s]", id, targetUrl)

	resp, err := ypdc.Client.Client.Get(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get [%s]: %w", ypdc.BaseUrl, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read body: %w", err)
	}

	var cards struct {
		Data []*model.Card `json:"data"`
	}
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return cards.Data[0], nil
}

func (ypdc YgoProDeckClient) SaveAllCards(_ *[]*model.Card) error {
	return fmt.Errorf("operation not supported")
}

func (ypdc YgoProDeckClient) SaveCard(_ *model.Card) error {
	return fmt.Errorf("operation not supported")
}

// Close is not necessary for the web client
func (ypdc YgoProDeckClient) Close() error { return nil }
