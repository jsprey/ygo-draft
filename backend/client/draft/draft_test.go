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
		dbMock := &mocks.DatabaseClient{}

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		require.ErrorIs(t, err, model.ErrorDraftIsNotAChallenge)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("fail as user is not receiver of challenge", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock.On("Exec", "UPDATE drafts\nSET status = 'running'\nWHERE id = 1;").Return(nil, assert.AnError)

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock.On("Exec", "UPDATE drafts\nSET status = 'running'\nWHERE id = 1;").Return(nil, nil)
		dbMock.On("Exec", "INSERT INTO draft_rounds (draft_id, round_number, status)\nVALUES (1, 1, 'preparation');").Return(nil, assert.AnError)

		client, err := NewDraftClient(dbMock)
		require.NoError(t, err)

		// when
		err = client.AcceptDraftChallenge(1, 3)

		// then
		require.ErrorIs(t, err, assert.AnError)
		mock.AssertExpectationsForObjects(t, dbMock)
	})
	t.Run("successfully accept ", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
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
		dbMock.On("Exec", "UPDATE drafts\nSET status = 'running'\nWHERE id = 1;").Return(nil, nil)
		dbMock.On("Exec", "INSERT INTO draft_rounds (draft_id, round_number, status)\nVALUES (1, 1, 'preparation');").Return(nil, nil)

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Exec", "INSERT INTO drafts (challenger_id, receiver_id, current_round_number, maximum_round_number, status, settings)\nVALUES (1, 2, 1, 10, 'pending', '{\"main_deck_draws\":1,\"main_deck_size\":2,\"extra_deck_draws\":3,\"extra_deck_size\":4,\"mode\":\"round\",\"mode_value\":10,\"sets\":[{\"set_name\":\"Test Set\",\"set_code\":\"TS\",\"set_rarity\":\"S\",\"set_rarity_code\":\"SS\"}]}');").Return(nil, assert.AnError)

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
	t.Run("fails when inserting new challenge", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'pending' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'pending' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Select", "SELECT d.id,\n       d.challenger_id,\n       d.receiver_id,\n       d.current_round_number,\n       d.maximum_round_number,\n       d.challenge_date,\n       d.winner_user_id,\n       d.status,\n       d.settings\nFROM drafts d\nWHERE (d.status = 'running' AND d.challenger_id = 1 AND d.receiver_id = 2)\n   OR (d.status = 'running' AND d.receiver_id = 1 AND d.challenger_id = 2)", mock.AnythingOfType("*[]model.Draft")).Return(nil)
		dbMock.On("Exec", "INSERT INTO drafts (challenger_id, receiver_id, current_round_number, maximum_round_number, status, settings)\nVALUES (1, 2, 1, 10, 'pending', '{\"main_deck_draws\":1,\"main_deck_size\":2,\"extra_deck_draws\":3,\"extra_deck_size\":4,\"mode\":\"round\",\"mode_value\":10,\"sets\":[{\"set_name\":\"Test Set\",\"set_code\":\"TS\",\"set_rarity\":\"S\",\"set_rarity_code\":\"SS\"}]}');").Return(nil, nil)

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock.On("Exec", "UPDATE drafts\nSET status = 'declined'\nWHERE id = 1;").Return(nil, assert.AnError)

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock.On("Exec", "UPDATE drafts\nSET status = 'declined'\nWHERE id = 1;").Return(nil, nil)

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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}

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
		draft, err := client.GetDraft(1, 1)

		// then
		require.NoError(t, err)
		require.Equal(t, expectedDraft, draft)
	})
}

func Test_draftClient_GetDraftsWithStatus(t *testing.T) {
	t.Run("fails to select query", func(t *testing.T) {
		// given
		dbMock := &mocks.DatabaseClient{}
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
		dbMock := &mocks.DatabaseClient{}

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
	type fields struct {
		Client         model.DatabaseClient
		QueryTemplater model.DraftQueryGenerator
	}
	type args struct {
		draftID            int
		surrenderingUserID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := draftClient{
				Client:         tt.fields.Client,
				QueryTemplater: tt.fields.QueryTemplater,
			}
			if err := d.SurrenderRunningDraft(tt.args.draftID, tt.args.surrenderingUserID); (err != nil) != tt.wantErr {
				t.Errorf("SurrenderRunningDraft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_draftClient_UserHaveDraftWithStatus(t *testing.T) {
	type fields struct {
		Client         model.DatabaseClient
		QueryTemplater model.DraftQueryGenerator
	}
	type args struct {
		user1ID int
		user2ID int
		status  model.DraftStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := draftClient{
				Client:         tt.fields.Client,
				QueryTemplater: tt.fields.QueryTemplater,
			}
			got, err := d.UserHaveDraftWithStatus(tt.args.user1ID, tt.args.user2ID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserHaveDraftWithStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserHaveDraftWithStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
