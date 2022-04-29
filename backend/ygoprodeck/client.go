package ygoprodeck

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"ygodraft/backend/model"
)

type YgoProDeckClient struct {
}

func NewYgoProDeckClient() *YgoProDeckClient {
	return &YgoProDeckClient{}
}

type response struct {
	Data []model.Card `json:"data"`
}

func (ypdc YgoProDeckClient) GetAllCards() (*[]model.Card, error) {
	resp, err := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		return nil, err
	}

	var cards response
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	return &cards.Data, nil
}
