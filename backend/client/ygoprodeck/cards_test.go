package ygoprodeck_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/client/ygoprodeck"
	"ygodraft/backend/client/ygoprodeck/mocks"
	"ygodraft/backend/model"
)

func TestYgoProDeckClient_GetAllCards(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		// given
		card1 := model.Card{Name: "mycard1"}
		card2 := model.Card{Name: "mycard2"}
		card3 := model.Card{Name: "mycard3"}

		limiter := &mocks.JsonWebClient{}
		limiter.On("GetJsonFromTarget", "/cardinfo.php", mock.Anything).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			expectedCards, ok := arg1.(*ygoprodeck.GetCardInfoResponse)
			if !ok {
				t.FailNow()
			}

			expectedCards.Data = []*model.Card{&card1, &card2, &card3}
		}).Return(nil)
		ygoClient := ygoprodeck.YgoProDeckClient{Client: limiter}

		// when
		cards, err := ygoClient.GetAllCards()

		// then
		require.NoError(t, err)
		assert.Len(t, *cards, 3)
		assert.Equal(t, "mycard1", (*cards)[0].Name)
		assert.Equal(t, "mycard2", (*cards)[1].Name)
		assert.Equal(t, "mycard3", (*cards)[2].Name)
	})

	t.Run("errors on get request", func(t *testing.T) {
		// given
		limiter := &mocks.JsonWebClient{}
		limiter.On("GetJsonFromTarget", "/cardinfo.php", mock.Anything).Return(assert.AnError)
		ygoClient := ygoprodeck.YgoProDeckClient{Client: limiter}

		// when
		cards, err := ygoClient.GetAllCards()

		// then
		require.Error(t, err)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, cards)
	})
}

func TestYgoProDeckClient_GetCard(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		// given
		expectedCard := model.Card{ID: 12345, Name: "mycard1"}

		limiter := &mocks.JsonWebClient{}
		limiter.On("GetJsonFromTarget", fmt.Sprintf("/cardinfo.php?id=%d", expectedCard.ID), mock.Anything).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			expectedCards, ok := arg1.(*ygoprodeck.GetCardInfoResponse)
			if !ok {
				t.FailNow()
			}

			expectedCards.Data = []*model.Card{&expectedCard}
		}).Return(nil)
		ygoClient := ygoprodeck.YgoProDeckClient{Client: limiter}

		// when
		acutalCard, err := ygoClient.GetCard(expectedCard.ID)

		// then
		require.NoError(t, err)
		assert.Equal(t, "mycard1", acutalCard.Name)
		assert.Equal(t, 12345, acutalCard.ID)
	})

	t.Run("errors on get request", func(t *testing.T) {
		// given
		expectedCard := model.Card{ID: 12345, Name: "mycard1"}
		limiter := &mocks.JsonWebClient{}
		limiter.On("GetJsonFromTarget", fmt.Sprintf("/cardinfo.php?id=%d", expectedCard.ID), mock.Anything).Return(assert.AnError)
		ygoClient := ygoprodeck.YgoProDeckClient{Client: limiter}

		// when
		cards, err := ygoClient.GetCard(expectedCard.ID)

		// then
		require.Error(t, err)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, cards)
	})
}
