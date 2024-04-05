package draft

import (
	"fmt"
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

	updateDraftQuery, err := d.QueryTemplater.UpdateDraft(draftID, model.DraftStatusRunning)
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

	updateDraftQuery, err := d.QueryTemplater.UpdateDraft(draftID, model.DraftStatusDeclined)
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

func (d draftClient) SurrenderRunningDraft(draftID int, surrenderingUserID int) error {
	//TODO implement me
	panic("implement me")
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
