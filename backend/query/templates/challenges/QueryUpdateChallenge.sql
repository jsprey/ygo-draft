-- Update a challenge
UPDATE public.draft_challenge
SET status = {{.Status}}
WHERE id = {{.ChallengeID}};