INSERT INTO draft_rounds (draft_id, round_number, status, winner_user_id)
VALUES ({{.DraftID}}, {{.RoundNumber}}, {{.Status}}, -1);