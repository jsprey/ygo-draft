package ygoprodeck

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ygodraft/backend/model"
)

// GetCardInfoResponse contains the data send when requesting the cardinfo endpoint.
type GetCardInfoResponse struct {
	Data []*model.Card `json:"data"`
}

// GetAllCards retrieves all cards from the ygo pro deck api.
func (ypdc YgoProDeckClient) GetAllCards() (*[]*model.Card, error) {
	targetUrl := fmt.Sprintf("%s/cardinfo.php", ypdc.BaseUrl)
	logrus.Debugf("YgoProDeckClient -> GetAllCards -> Requesting [%s]", targetUrl)

	var cards GetCardInfoResponse
	err := ypdc.Client.GetJsonFromTarget(targetUrl, &cards)
	if err != nil {
		return nil, err
	}

	return &cards.Data, nil
}

// GetCard retrieves a api with the given id from the ygo pro deck api.
func (ypdc YgoProDeckClient) GetCard(id int) (*model.Card, error) {
	targetUrl := fmt.Sprintf("%s/cardinfo.php?id=%d", ypdc.BaseUrl, id)
	logrus.Debugf("YgoProDeckClient -> GetCard [%d] -> Requesting [%s]", id, targetUrl)

	var cards GetCardInfoResponse
	err := ypdc.Client.GetJsonFromTarget(targetUrl, &cards)
	if err != nil {
		return nil, err
	}

	return cards.Data[0], nil
}
