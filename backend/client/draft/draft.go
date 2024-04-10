package draft

import (
	"fmt"
	"strings"
	"ygodraft/backend/model"
	"ygodraft/backend/query"
)

type draftClient struct {
	Client         model.DatabaseClient
	QueryTemplater model.DraftQueryGenerator
}

func (d draftClient) CreateDraftChallenge(challengerID int, receiverID int, settings model.DraftSettings) error {
	// check for present challenge
	hasDraftChallenge, err := d.UserHaveDraftWithStatus(challengerID, receiverID, model.DraftStatusPending)
	if err != nil {
		return fmt.Errorf("failed to check for existing draft challenge: %w", err)
	}

	if hasDraftChallenge {
		return model.ErrorUserAlreadyChallenged
	}

	// check for present running draft
	hasRunningDraft, err := d.UserHaveDraftWithStatus(challengerID, receiverID, model.DraftStatusRunning)
	if err != nil {
		return fmt.Errorf("failed to check for existing running draft: %w", err)
	}

	if hasRunningDraft {
		return model.ErrorUserAlreadyRunningDraft
	}

	// create new draft challenge
	insertDraftQuery, err := d.QueryTemplater.InsertDraft(challengerID, receiverID, settings)
	if err != nil {
		return err
	}

	_, err = d.Client.Exec(insertDraftQuery)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (d draftClient) UserHaveDraftWithStatus(user1ID int, user2ID int, status model.DraftStatus) (bool, error) {
	selectDraftQuery, err := d.QueryTemplater.SelectDraftsWithUsersAndStatus(user1ID, user2ID, status)
	if err != nil {
		return false, fmt.Errorf("failed to template query: %w", err)
	}

	var drafts []model.Draft
	err = d.Client.Select(selectDraftQuery, &drafts)
	if err != nil {
		return false, fmt.Errorf("failed to select: %w", err)
	}

	if drafts == nil {
		return false, nil
	}

	return len(drafts) > 0, nil
}

func (d draftClient) AcceptDraftChallenge(draftID int, userID int) error {
	draft, err := d.GetDraft(draftID, userID)
	if err != nil {
		return fmt.Errorf("failed to get draft: %w", err)
	}

	if draft.Status == model.DraftStatusRunning {
		return nil
	}

	// check if draft is challenge at all
	if draft.Status != model.DraftStatusPending {
		return model.ErrorDraftIsNotAChallenge
	}

	// check if user id is receiver
	if userID != draft.ReceiverID {
		return model.ErrorOnlyReceivingPartyCanAcceptChallenge
	}

	updateDraftQuery, err := d.QueryTemplater.UpdateDraft(draftID, draft.CurrentRoundNumber, draft.WinnerUserID, model.DraftStatusRunning)
	if err != nil {
		return fmt.Errorf("failed to template [UpdateDraft] query: %w", err)
	}

	_, err = d.Client.Exec(updateDraftQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateDraft] query: %w", err)
	}

	// create first round
	insertDraftRoundQuery, err := d.QueryTemplater.InsertDraftRound(draftID, 1, model.DraftRoundStatusPreparation)
	if err != nil {
		return fmt.Errorf("failed to template [InsertDraftRound] query: %w", err)
	}

	_, err = d.Client.Exec(insertDraftRoundQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [InsertDraftRound] query: %w", err)
	}

	return nil
}

func (d draftClient) DeclineDraftChallenge(draftID int, userID int) error {
	draft, err := d.GetDraft(draftID, userID)
	if err != nil {
		return fmt.Errorf("failed to get draft: %w", err)
	}

	if draft.Status == model.DraftStatusDeclined {
		return nil
	}

	// check if draft is challenge at all
	if draft.Status != model.DraftStatusPending {
		return model.ErrorDraftIsNotAChallenge
	}

	// check if user id is receiver
	if userID != draft.ReceiverID {
		return model.ErrorOnlyReceivingPartyCanDeclineChallenge
	}

	updateDraftQuery, err := d.QueryTemplater.UpdateDraft(draftID, draft.CurrentRoundNumber, draft.WinnerUserID, model.DraftStatusDeclined)
	if err != nil {
		return fmt.Errorf("failed to template [UpdateDraft] query: %w", err)
	}

	_, err = d.Client.Exec(updateDraftQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateDraft] query: %w", err)
	}

	return nil
}

func (d draftClient) GetDraftsWithStatus(userID int, status model.DraftStatus) ([]model.Draft, error) {
	selectDraftQuery, err := d.QueryTemplater.SelectDraftsWithStatus(userID, status)
	if err != nil {
		return nil, err
	}

	var drafts []model.Draft
	err = d.Client.Select(selectDraftQuery, &drafts)
	if err != nil {
		return []model.Draft{}, fmt.Errorf("failed to select: %w", err)
	}

	if drafts == nil {
		return []model.Draft{}, nil
	}

	return drafts, nil
}

func (d draftClient) GetDraft(draftID int, userID int) (model.Draft, error) {
	getDraftQuery, err := d.QueryTemplater.SelectDraft(draftID)
	if err != nil {
		return model.Draft{}, fmt.Errorf("failed to template [SelectDraft] query: %w", err)
	}

	var drafts []model.Draft
	err = d.Client.Select(getDraftQuery, &drafts)
	if err != nil {
		return model.Draft{}, fmt.Errorf("failed to select query [SelectDraft]: %w", err)
	}

	if drafts == nil {
		drafts = []model.Draft{}
	}

	if len(drafts) == 0 {
		return model.Draft{}, model.ErrorDraftDoesNotExist.WithParam(string(rune(draftID)))
	}

	draft := drafts[0]
	if draft.ReceiverID != userID && draft.ChallengerID != userID {
		return model.Draft{}, model.ErrorUserIsNotParticipatingInDraft
	}

	return drafts[0], nil
}

func (d draftClient) GetDraftRounds(draftID int) ([]model.DraftRound, error) {
	draftRoundsQuery, err := d.QueryTemplater.SelectDraftRounds(draftID)
	if err != nil {
		return []model.DraftRound{}, fmt.Errorf("failed to template [SelectDraftRounds]: %w", err)
	}

	var draftRounds []model.DraftRound
	err = d.Client.Select(draftRoundsQuery, &draftRounds)
	if err != nil {
		return []model.DraftRound{}, fmt.Errorf("failed to select query [SelectDraftRounds]: %w", err)
	}

	if draftRounds == nil {
		draftRounds = []model.DraftRound{}
	}

	return draftRounds, nil
}

func (d draftClient) SurrenderRunningDraft(draftID int, surrenderingUserID int) error {
	//TODO implement me
	panic("implement me")
}

func (d draftClient) GetDraftRound(roundID int, userID int) (model.DraftRound, error) {
	roundQuery, err := d.QueryTemplater.SelectDraftRound(roundID)
	if err != nil {
		return model.DraftRound{}, fmt.Errorf("failed to template query [SelectDraftRound]: %w", err)
	}

	var draftRounds []model.DraftRound
	err = d.Client.Select(roundQuery, &draftRounds)
	if err != nil {
		return model.DraftRound{}, fmt.Errorf("failed to template query [SelectDraftRound]: %w", err)
	}

	if draftRounds == nil || len(draftRounds) == 0 {
		return model.DraftRound{}, model.ErrorDraftRoundDoesNotExist.WithParam(fmt.Sprintf("%d", roundID))
	}

	// check users access to the draft
	draftRound := draftRounds[0]
	_, err = d.GetDraft(draftRound.DraftID, userID)
	if err != nil {
		return model.DraftRound{}, fmt.Errorf("failed to get draft: %w", err)
	}

	return draftRound, nil
}

func (d draftClient) GetDraftRoundDeck(roundID int, userID int) ([]string, error) {
	_, err := d.GetDraftRound(roundID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get draft round: %w", err)
	}

	deckQuery, err := d.QueryTemplater.SelectDraftRoundDeck(roundID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to template query [SelectDraftRoundDeck]: %w", err)
	}

	var deckList []model.DraftRoundDeck
	err = d.Client.Select(deckQuery, &deckList)
	if err != nil {
		return nil, fmt.Errorf("failed to select query [SelectDraftRoundDeck]: %w", err)
	}

	if deckList == nil || len(deckList) == 0 {
		return []string{}, model.ErrorDraftRoundDeckDoesNotExist.WithParam(roundID, userID)
	}

	var deckListRaw = deckList[0].Deck
	return strings.Split(deckListRaw, ","), nil
}

func (d draftClient) SubmitDraftDeck(roundID int, userID int, deck []string) error {
	userInsertRoundDeckQuery, err := d.QueryTemplater.InsertDraftRoundDeck(roundID, userID, deck)
	if err != nil {
		return fmt.Errorf("failed to template query [InsertDraftRoundDeck]: %w", err)
	}

	_, err = d.Client.Exec(userInsertRoundDeckQuery)
	if err != nil {
		return fmt.Errorf("failed to exec query [InsertDraftRoundDeck]: %w", err)
	}

	round, err := d.GetDraftRound(roundID, userID)
	if err != nil {
		return fmt.Errorf("failed to get draft round: %w", err)
	}

	draft, err := d.GetDraft(round.DraftID, userID)
	if err != nil {
		return fmt.Errorf("failed to get draft: %w", err)
	}

	enemyID := draft.ChallengerID
	if userID == draft.ChallengerID {
		enemyID = draft.ReceiverID
	}

	_, err = d.GetDraftRoundDeck(roundID, enemyID)
	if model.IsErrorCustom(err, model.ErrorDraftRoundDeckDoesNotExist) {
		// enemy must still draft his deck
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get draft round deck: %w", err)
	}

	// change status of round to fight
	changeRoundStatus, err := d.QueryTemplater.UpdateDraftRound(roundID, -1, model.DraftRoundStatusFighting)
	if err != nil {
		return fmt.Errorf("failed to template [UpdateDraftRound] query: %w", err)
	}

	_, err = d.Client.Exec(changeRoundStatus)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateDraftRound] query: %w", err)
	}

	return nil
}

func (d draftClient) SetWinnerForDraftRound(roundID int, winnerUser int) error {
	// change status of round to finished
	changeRoundStatus, err := d.QueryTemplater.UpdateDraftRound(roundID, winnerUser, model.DraftRoundStatusFinished)
	if err != nil {
		return fmt.Errorf("failed to template [UpdateDraftRound] query: %w", err)
	}

	_, err = d.Client.Exec(changeRoundStatus)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateDraftRound] query: %w", err)
	}

	draftRound, err := d.GetDraftRound(roundID, winnerUser)
	if err != nil {
		return fmt.Errorf("failed to get draft round: %w", err)
	}

	currentDraft, err := d.GetDraft(draftRound.DraftID, winnerUser)
	if err != nil {
		return fmt.Errorf("failed to get draft: %w", err)
	}

	// create next round
	insertDraftRoundQuery, err := d.QueryTemplater.InsertDraftRound(draftRound.DraftID, currentDraft.CurrentRoundNumber+1, model.DraftRoundStatusPreparation)
	if err != nil {
		return fmt.Errorf("failed to template [InsertDraftRound] query: %w", err)
	}

	_, err = d.Client.Exec(insertDraftRoundQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [InsertDraftRound] query: %w", err)
	}

	// update current round in draft
	draftUpdateQuery, err := d.QueryTemplater.UpdateDraft(currentDraft.ID, currentDraft.CurrentRoundNumber+1, currentDraft.WinnerUserID, currentDraft.Status)
	if err != nil {
		return fmt.Errorf("failed to template query [UpdateDraft]: %w", err)
	}

	_, err = d.Client.Exec(draftUpdateQuery)
	if err != nil {
		return fmt.Errorf("failed to exec [UpdateDraft] query: %w", err)
	}

	return nil
}

func NewDraftClient(dbClient model.DatabaseClient) (*draftClient, error) {
	queryTemplater, err := query.NewSqlQueryTemplater()
	if err != nil {
		return nil, fmt.Errorf("failed to create new sql query templater: %w", err)
	}

	return &draftClient{
		Client:         dbClient,
		QueryTemplater: queryTemplater,
	}, nil
}
