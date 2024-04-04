UPDATE draft_rounds
SET status = {{.Status}}{{if .WinnerUserID}},
    winnerUserID = {{.WinnerUserID}}{{end}}
WHERE id = {{.RoundID}};
