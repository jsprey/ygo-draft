package cache_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"ygodraft/backend/client/cache"
	"ygodraft/backend/model"
	mocks2 "ygodraft/backend/model/mocks"
)

func getCacheWithMocks() (*cache.YgoCache, *mocks2.QueryGenerator, *mocks2.DatabaseClient) {
	myCache := &cache.YgoCache{}

	queryGenMock := &mocks2.QueryGenerator{}
	myCache.QueryTemplater = queryGenMock

	dbMock := &mocks2.DatabaseClient{}
	myCache.Client = dbMock

	return myCache, queryGenMock, dbMock
}

func TestYgoCache_GetAllCards(t *testing.T) {
	t.Run("successfully get all cards", func(t *testing.T) {
		// given
		expectedCards := []*model.Card{
			{ID: 123},
			{ID: 456},
			{ID: 789},
		}

		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectAllCards").Return("myquery", nil)

		var argsCars []*model.Card
		dbMock.On("Select", "myquery", &argsCars).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			cardsInArgs, ok := arg1.(*[]*model.Card)
			if !ok {
				t.FailNow()
			}

			*cardsInArgs = append(*cardsInArgs, expectedCards[0])
			*cardsInArgs = append(*cardsInArgs, expectedCards[1])
			*cardsInArgs = append(*cardsInArgs, expectedCards[2])
		}).Return(nil)

		// when
		allCards, err := myCache.GetAllCards()

		// then
		require.NoError(t, err)
		require.NotNil(t, allCards)
		assert.Len(t, *allCards, 3)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("fails for query generation", func(t *testing.T) {
		// given
		myCache, queryGenerator, _ := getCacheWithMocks()
		queryGenerator.On("SelectAllCards").Return("myquery", assert.AnError)

		// when
		_, err := myCache.GetAllCards()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator)
	})

	t.Run("fails for query", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectAllCards").Return("myquery", nil)
		dbMock.On("Select", "myquery", mock.Anything).Return(assert.AnError)

		// when
		_, err := myCache.GetAllCards()

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})
}

func TestYgoCache_GetCard(t *testing.T) {
	t.Run("successfully get cards by id", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectCardByID", 444).Return("myquery", nil)

		var argsCars []*model.Card
		dbMock.On("Select", "myquery", &argsCars).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			cardsInArgs, ok := arg1.(*[]*model.Card)
			if !ok {
				t.FailNow()
			}

			*cardsInArgs = append(*cardsInArgs, &model.Card{ID: 444})
		}).Return(nil)

		// when
		card, err := myCache.GetCard(444)

		// then
		require.NoError(t, err)
		require.NotNil(t, card)
		assert.Equal(t, 444, card.ID)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("card does not exist", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectCardByID", 11).Return("myquery", nil)

		var argsCards []*model.Card
		dbMock.On("Select", "myquery", &argsCards).Return(nil)

		// when
		_, err := myCache.GetCard(11)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, cache.ErrorCardDoesNotExist)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("fails for query generation", func(t *testing.T) {
		// given
		myCache, queryGenerator, _ := getCacheWithMocks()
		queryGenerator.On("SelectCardByID", 444).Return("myquery", assert.AnError)

		// when
		_, err := myCache.GetCard(444)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator)
	})

	t.Run("fails for query", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectCardByID", 444).Return("myquery", nil)
		dbMock.On("Select", "myquery", mock.Anything).Return(assert.AnError)

		// when
		_, err := myCache.GetCard(444)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})
}

func TestYgoCache_GetSet(t *testing.T) {
	t.Run("successfully get set by code", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectSetByCode", "AABBCC").Return("myquery", nil)

		var argSets []*model.CardSet
		dbMock.On("Select", "myquery", &argSets).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			setInArgs, ok := arg1.(*[]*model.CardSet)
			if !ok {
				t.FailNow()
			}

			*setInArgs = append(*setInArgs, &model.CardSet{
				SetName:       "AA-BB-CC",
				SetCode:       "AABBCC",
				SetRarity:     "selten",
				SetRarityCode: "d",
			})
		}).Return(nil)

		// when
		set, err := myCache.GetSet("AABBCC")

		// then
		require.NoError(t, err)
		require.NotNil(t, set)
		assert.Equal(t, "AABBCC", set.SetCode)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("set does not exist", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectSetByCode", "AABBCC").Return("myquery", nil)

		var argSets []*model.CardSet
		dbMock.On("Select", "myquery", &argSets).Return(nil)

		// when
		_, err := myCache.GetSet("AABBCC")

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, cache.ErrorSetDoesNotExist)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("fails for query generation", func(t *testing.T) {
		// given
		myCache, queryGenerator, _ := getCacheWithMocks()
		queryGenerator.On("SelectSetByCode", "AABBCC").Return("", assert.AnError)

		// when
		_, err := myCache.GetSet("AABBCC")

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator)
	})

	t.Run("fails for query", func(t *testing.T) {
		// given
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("SelectSetByCode", "AABBCC").Return("myquery", nil)
		dbMock.On("Select", "myquery", mock.Anything).Return(assert.AnError)

		// when
		_, err := myCache.GetSet("AABBCC")

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})
}

func TestYgoCache_SaveAllCards(t *testing.T) {
	t.Run("successfully insert", func(t *testing.T) {
		// given
		card111 := &model.Card{
			ID:   111,
			Name: "name 1",
		}
		card222 := &model.Card{
			ID:   222,
			Name: "name 2",
		}
		card333 := &model.Card{
			ID:   333,
			Name: "name 3",
		}
		allCards := []*model.Card{card111, card222, card333}

		myCache, queryGenerator, dbMock := getCacheWithMocks()

		queryGenerator.On("InsertCard", card222).Return("myinsert222", nil)
		queryGenerator.On("InsertCard", card333).Return("myinsert333", nil)
		dbMock.On("Exec", "myinsert222").Return(nil, nil)
		dbMock.On("Exec", "myinsert333").Return(nil, nil)

		queryGenerator.On("SelectCardByID", 111).Return("myquery111", nil)
		queryGenerator.On("SelectCardByID", 222).Return("myquery222", nil)
		queryGenerator.On("SelectCardByID", 333).Return("myquery333", nil)

		var argsCars111 []*model.Card
		dbMock.On("Select", "myquery111", &argsCars111).Run(func(args mock.Arguments) {
			arg1 := args.Get(1)
			cardsInArgs, ok := arg1.(*[]*model.Card)
			if !ok {
				t.FailNow()
			}

			*cardsInArgs = append(*cardsInArgs, card111)
		}).Return(nil)
		var argsCards222 []*model.Card
		dbMock.On("Select", "myquery222", &argsCards222).Return(nil)
		var argsCards333 []*model.Card
		dbMock.On("Select", "myquery333", &argsCards333).Return(nil)

		// when
		err := myCache.SaveAllCards(&allCards)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})
}

func TestYgoCache_SaveCard(t *testing.T) {
	t.Run("successfully insert", func(t *testing.T) {
		// given
		insertCard := &model.Card{
			ID:        555,
			Name:      "name",
			Type:      "type",
			Desc:      "desc",
			Atk:       100,
			Def:       200,
			Level:     300,
			Race:      "myRace",
			Attribute: "myAttribute",
		}

		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("InsertCard", insertCard).Return("myquery", nil)
		dbMock.On("Exec", "myquery").Return(nil, nil)

		// when
		err := myCache.SaveCard(insertCard)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("error on query generation", func(t *testing.T) {
		// given
		myCard := &model.Card{ID: 123123}
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("InsertCard", myCard).Return("myquery", assert.AnError)

		// when
		err := myCache.SaveCard(myCard)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})

	t.Run("fails for exec", func(t *testing.T) {
		// given
		myCard := &model.Card{ID: 123123}
		myCache, queryGenerator, dbMock := getCacheWithMocks()
		queryGenerator.On("InsertCard", myCard).Return("myquery", nil)
		dbMock.On("Exec", "myquery").Return(nil, assert.AnError)

		// when
		err := myCache.SaveCard(myCard)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, queryGenerator, dbMock)
	})
}
