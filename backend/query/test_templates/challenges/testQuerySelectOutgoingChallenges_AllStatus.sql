-- Select all outgoing challenges
SELECT dc.id, dc.challenger_id, dc.receiver_id, dc.challenge_date, dc.status, dc.settings
FROM draft_challenge as dc
WHERE dc.challenger_id = 0;