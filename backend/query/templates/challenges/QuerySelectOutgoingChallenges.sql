-- Select all outgoing challenges
SELECT dc.id, dc.challenger_id, dc.receiver_id, dc.challenge_date, dc.status, dc.settings
FROM draft_challenges as dc
WHERE dc.challenger_id = {{.ChallengerID}}{{if.Status}} AND dc.status = {{.Status}}{{end}};