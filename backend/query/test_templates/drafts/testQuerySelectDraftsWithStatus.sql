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
WHERE (d.status = 'running' AND d.challenger_id = 4)
   OR (d.status = 'running' AND d.receiver_id = 4)