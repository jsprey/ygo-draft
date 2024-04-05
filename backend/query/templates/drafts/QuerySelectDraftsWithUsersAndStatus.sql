SELECT d.id,
       d.challenger_id,
       d.receiver_id,
       d.current_round_number,
       d.maximum_round_number,
       d.challenge_date,
       d.winner_user_id,
       d.status,
       d.settings
FROM drafts d
WHERE (d.status = {{.Status}} AND d.challenger_id = {{.UserID}} AND d.receiver_id = {{.User2ID}})
   OR (d.status = {{.Status}} AND d.receiver_id = {{.UserID}} AND d.challenger_id = {{.User2ID}})