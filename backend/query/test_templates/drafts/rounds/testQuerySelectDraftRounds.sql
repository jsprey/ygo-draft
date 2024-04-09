SELECT d.id,
       d.draft_id,
       d.round_number,
       d.winner_user_id,
       d.status
FROM ygodraft.public.draft_rounds d
WHERE d.draft_id = 4;