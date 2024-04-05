UPDATE draft_rounds
SET status = {{.Status}}{{if .WinnerUserID}},
    winner_user_id = {{.WinnerUserID}}{{end}}
WHERE id = {{.RoundID}};
