-- Insert a new challenge
INSERT INTO public.draft_challenges (challenger_id, receiver_id, status, settings)
VALUES ({{.ChallengerID}}, {{.ReceiverID}}, {{.Status}}, {{.Settings}})