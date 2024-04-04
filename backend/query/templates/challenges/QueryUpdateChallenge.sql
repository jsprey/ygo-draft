-- Update a challenge
UPDATE public.draft_challenges
SET status = {{.Status}}
WHERE id = {{.ChallengeID}};