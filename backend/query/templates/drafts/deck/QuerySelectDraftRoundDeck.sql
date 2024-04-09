SELECT d.id,
       d.user_id,
       d.round_id,
       d.deck
FROM ygodraft.public.draft_deck d
WHERE d.id = {{.UserID}} AND d.round_id = {{.RoundID}};