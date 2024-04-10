SELECT d.id,
       d.user_id,
       d.round_id,
       d.deck
FROM ygodraft.public.draft_deck d
WHERE d.user_id = 4 AND d.round_id = 5;