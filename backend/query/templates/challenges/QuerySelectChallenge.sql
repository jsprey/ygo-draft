-- Select a specific challenge
SELECT dc.id, dc.challenger_id, dc.receiver_id, dc.challenge_date, dc.status, dc.settings
FROM draft_challenges as dc
WHERE dc.id = {{.ChallengeID}};