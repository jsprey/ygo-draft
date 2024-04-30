package draft

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"ygodraft/backend/model"
	"ygodraft/backend/model/mocks"
)

func TestNewDraftClient(t *testing.T) {
	t.Run("successfully create new draft client", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		// when
		client, err := NewDraftClient(dbMock)

		// then
		require.NoError(t, err)
		require.Implements(t, (*model.DraftClient)(nil), client)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
}

func Test_draftClient_AcceptDraftChallenge(t *testing.T) {
	t.Run("fail as cannot get draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(0, 0)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as draft is not a challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusDeclined,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, model.ErrorDraftIsNotAChallenge)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("ignore and do nothing when challenge is currently a running draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusRunning,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user is not receiver of challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 2)

		// then
		require.ErrorIs(t, err, model.ErrorOnlyReceivingPartyCanAcceptChallenge)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail to exec update draft query", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'running',\n    winner_user_id = 0,\n    current_round_number = 0\nWHERE id = 1;").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail to exec insert draft round query", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'running',\n    winner_user_id = 0,\n    current_round_number = 0\nWHERE id = 1;").Return(nil, nil)
		dbMock.On("Exec", "INSERT INTO draft_rounds (draft_id, round_number, status, winner_user_id)\nVALUES (1, 1, 'preparation', -1);").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("successfully accept", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'running',\n    winner_user_id = 0,\n    current_round_number = 0\nWHERE id = 1;").Return(nil, nil)
		dbMock.On("Exec", "INSERT INTO draft_rounds (draft_id, round_number, status, winner_user_id)\nVALUES (1, 1, 'preparation', -1);").Return(nil, nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
}

func Test_draftClient_CreateDraftChallenge(t *testing.T) {
	t.Run("fail as cannot get draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user already has challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.ErrorIs(t, err, model.ErrorUserAlreadyChallenged)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user already has challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user already has a running draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusRunning,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.ErrorIs(t, err, model.ErrorUserAlreadyRunningDraft)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fails when inserting new challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Exec", "INSERT INTO drafts (challenger_id, receiver_id, current_round_number, maximum_round_number, winner_user_id, status, settings)\nVALUES (1, 2, 1, 10, -1, 'pending', '{\"main_deck_draws\":1,\"main_deck_size\":2,\"extra_deck_draws\":3,\"extra_deck_size\":4,\"mode\":\"rounds\",\"mode_value\":10,\"sets\":[{\"set_name\":\"Test Set\",\"set_code\":\"TS\",\"set_rarity\":\"S\",\"set_rarity_code\":\"SS\"}]}');").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("successfully insert new challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Exec", "INSERT INTO drafts (challenger_id, receiver_id, current_round_number, maximum_round_number, winner_user_id, status, settings)\nVALUES (1, 2, 1, 10, -1, 'pending', '{\"main_deck_draws\":1,\"main_deck_size\":2,\"extra_deck_draws\":3,\"extra_deck_size\":4,\"mode\":\"rounds\",\"mode_value\":10,\"sets\":[{\"set_name\":\"Test Set\",\"set_code\":\"TS\",\"set_rarity\":\"S\",\"set_rarity_code\":\"SS\"}]}');").Return(nil, nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		settings := model.DraftSettings{
			MainDeckDraws:  1,
			MainDeckSize:   2,
			ExtraDeckDraws: 3,
			ExtraDeckSize:  4,
			Mode:           model.DraftModeRounds,
			ModeValue:      10,
			Sets:           []model.CardSet{{SetName: "Test Set", SetCode: "TS", SetRarity: "S", SetRarityCode: "SS"}},
		}

		// when
		err = client.CreateDraftChallenge(1, 2, settings)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
}

func Test_draftClient_DeclineDraftChallenge(t *testing.T) {
	t.Run("fail as cannot get draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.DeclineDraftChallenge(0, 0)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as draft is not a challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusRunning,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.DeclineDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, model.ErrorDraftIsNotAChallenge)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user is not receiver of challenge", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.DeclineDraftChallenge(1, 2)

		// then
		require.ErrorIs(t, err, model.ErrorOnlyReceivingPartyCanDeclineChallenge)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail to exec update draft query", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'declined',\n    winner_user_id = 0,\n    current_round_number = 0\nWHERE id = 1;").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.DeclineDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("successfully accept ", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", mock.AnythingOfType("string"), mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)

			draft := model.Draft{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			}

			*drafts = append(*drafts, draft)
		}).Return(nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'declined',\n    winner_user_id = 0,\n    current_round_number = 0\nWHERE id = 1;").Return(nil, nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.DeclineDraftChallenge(1, 3)

		// then
		require.NoError(t, err)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
}

func Test_draftClient_GetDraft(t *testing.T) {
	t.Run("fails to select", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		_, err = client.GetDraft(0, 0)

		// then
		require.ErrorIs(t, err, assert.AnError)

	})
	t.Run("fails as no draft exist", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 1;", mock.AnythingOfType("*[]model.Draft")).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		_, err = client.GetDraft(1, 1)

		// then
		require.ErrorIs(t, err, model.ErrorDraftDoesNotExist.WithParam(string(rune(1))))
	})
	t.Run("fails as no draft exist", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDraft := model.Draft{
			ID:            1,
			ChallengerID:  2,
			ReceiverID:    3,
			Status:        model.DraftStatusPending,
			Settings:      model.DraftSettings{},
			ChallengeDate: time.Now(),
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 1;", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = append(*drafts, expectedDraft)
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		draft, err := client.GetDraft(1, 2)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedDraft, draft)
	})
}

func Test_draftClient_GetDraftsWithStatus(t *testing.T) {
	t.Run("fails to select query", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'finished' AND d.challenger_id = 0)\n   OR (d.status = 'finished' AND d.receiver_id = 0)", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		_, err = client.GetDraftsWithStatus(0, model.DraftStatusFinished)

		// then
		require.ErrorIs(t, err, assert.AnError)

	})
	t.Run("fails as no draft exist", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			},
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1)\n   OR (d.status = 'pending' AND d.receiver_id = 1)", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		actualDrafts, err := client.GetDraftsWithStatus(1, model.DraftStatusPending)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedDrafts, actualDrafts)
	})
}

func Test_draftClient_SurrenderRunningDraft(t *testing.T) {
	t.Run("fails to select query", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.SurrenderRunningDraft(0, 4)

		// then
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("throw error when draft is not running", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:            1,
				ChallengerID:  2,
				ReceiverID:    3,
				Status:        model.DraftStatusPending,
				Settings:      model.DraftSettings{},
				ChallengeDate: time.Now(),
			},
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.SurrenderRunningDraft(0, 2)

		// then
		require.ErrorIs(t, err, model.ErrorDraftIsNotRunning)
	})
	t.Run("fail to update draft rounds", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:                 1,
				ChallengerID:       2,
				ReceiverID:         3,
				Status:             model.DraftStatusRunning,
				MaximumRoundNumber: 3,
				CurrentRoundNumber: 2,
				WinnerUserID:       -1,
				Settings:           model.DraftSettings{},
				ChallengeDate:      time.Now(),
			},
		}
		expectedDraftRounds := []model.DraftRound{
			{
				ID:           0,
				DraftID:      1,
				RoundNumber:  1,
				Status:       model.DraftRoundStatusFinished,
				WinnerUserID: 3,
			},
			{
				ID:           1,
				DraftID:      1,
				RoundNumber:  2,
				Status:       model.DraftRoundStatusFighting,
				WinnerUserID: -1,
			},
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.draft_id,\n       d.round_number,\n       d.winner_user_id,\n       d.status\nFROM ygodraft.public.draft_rounds d\nWHERE d.draft_id = 1;", mock.AnythingOfType("*[]model.DraftRound")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			rounds, ok := get.(*[]model.DraftRound)
			require.True(t, ok)
			*rounds = expectedDraftRounds
		}).Return(nil)
		dbMock.On("Exec", "UPDATE draft_rounds\nSET status = 'finished',\n    winner_user_id = 2\nWHERE id = 1;\n").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.SurrenderRunningDraft(0, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("fail to update draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:                 1,
				ChallengerID:       2,
				ReceiverID:         3,
				Status:             model.DraftStatusRunning,
				MaximumRoundNumber: 3,
				CurrentRoundNumber: 2,
				WinnerUserID:       -1,
				Settings:           model.DraftSettings{},
				ChallengeDate:      time.Now(),
			},
		}
		expectedDraftRounds := []model.DraftRound{
			{
				ID:           0,
				DraftID:      1,
				RoundNumber:  1,
				Status:       model.DraftRoundStatusFinished,
				WinnerUserID: 3,
			},
			{
				ID:           1,
				DraftID:      1,
				RoundNumber:  2,
				Status:       model.DraftRoundStatusFighting,
				WinnerUserID: -1,
			},
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.draft_id,\n       d.round_number,\n       d.winner_user_id,\n       d.status\nFROM ygodraft.public.draft_rounds d\nWHERE d.draft_id = 1;", mock.AnythingOfType("*[]model.DraftRound")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			rounds, ok := get.(*[]model.DraftRound)
			require.True(t, ok)
			*rounds = expectedDraftRounds
		}).Return(nil)
		dbMock.On("Exec", "UPDATE draft_rounds\nSET status = 'finished',\n    winner_user_id = 2\nWHERE id = 1;\n").Return(nil, nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'surrender',\n    winner_user_id = 2,\n    current_round_number = 2\nWHERE id = 1;").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.SurrenderRunningDraft(0, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("successfully surrender draft", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:                 1,
				ChallengerID:       2,
				ReceiverID:         3,
				Status:             model.DraftStatusRunning,
				MaximumRoundNumber: 3,
				CurrentRoundNumber: 2,
				WinnerUserID:       -1,
				Settings:           model.DraftSettings{},
				ChallengeDate:      time.Now(),
			},
		}
		expectedDraftRounds := []model.DraftRound{
			{
				ID:           0,
				DraftID:      1,
				RoundNumber:  1,
				Status:       model.DraftRoundStatusFinished,
				WinnerUserID: 3,
			},
			{
				ID:           1,
				DraftID:      1,
				RoundNumber:  2,
				Status:       model.DraftRoundStatusFighting,
				WinnerUserID: -1,
			},
		}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE d.id = 0;", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.draft_id,\n       d.round_number,\n       d.winner_user_id,\n       d.status\nFROM ygodraft.public.draft_rounds d\nWHERE d.draft_id = 1;", mock.AnythingOfType("*[]model.DraftRound")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			rounds, ok := get.(*[]model.DraftRound)
			require.True(t, ok)
			*rounds = expectedDraftRounds
		}).Return(nil)
		dbMock.On("Exec", "UPDATE draft_rounds\nSET status = 'finished',\n    winner_user_id = 2\nWHERE id = 1;\n").Return(nil, nil)
		dbMock.On("Exec", "UPDATE ygodraft.public.drafts\nSET status = 'surrender',\n    winner_user_id = 2,\n    current_round_number = 2\nWHERE id = 1;").Return(nil, nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.SurrenderRunningDraft(0, 3)

		// then
		require.NoError(t, err)
	})
}

func Test_draftClient_UserHaveDraftWithStatus(t *testing.T) {
	t.Run("fail to get drafts", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 3 AND d.receiver_id = 4)\n   OR (d.status = 'running' AND d.receiver_id = 3 AND d.challenger_id = 4)", mock.AnythingOfType("*[]model.Draft")).Return(assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		_, err = client.UserHaveDraftWithStatus(3, 4, model.DraftStatusRunning)

		// then
		require.ErrorIs(t, err, assert.AnError)
	})
	t.Run("successfully get draft with status", func(t *testing.T) {
		// given
		dbMock := mocks.NewDatabaseClient(t)

		expectedDrafts := []model.Draft{
			{
				ID:                 1,
				ChallengerID:       3,
				ReceiverID:         4,
				Status:             model.DraftStatusRunning,
				MaximumRoundNumber: 3,
				CurrentRoundNumber: 2,
				WinnerUserID:       -1,
				Settings:           model.DraftSettings{},
				ChallengeDate:      time.Now(),
			},
		}

		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 3 AND d.receiver_id = 4)\n   OR (d.status = 'running' AND d.receiver_id = 3 AND d.challenger_id = 4)", mock.AnythingOfType("*[]model.Draft")).Run(func(args mock.Arguments) {
			get := args.Get(1)
			drafts, ok := get.(*[]model.Draft)
			require.True(t, ok)
			*drafts = expectedDrafts
		}).Return(nil)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		_, err = client.UserHaveDraftWithStatus(3, 4, model.DraftStatusRunning)

		// then
		require.NoError(t, err)
	})
}
