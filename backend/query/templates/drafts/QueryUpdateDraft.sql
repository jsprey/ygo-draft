UPDATE ygodraft.public.drafts
SET status = {{.Status}},
    winner_user_id = {{.WinnerID}},
    current_round_number = {{.CurrentRoundNumber}}
WHERE id = {{.DraftID}};