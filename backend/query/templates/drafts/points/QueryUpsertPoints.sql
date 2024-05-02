INSERT INTO ygodraft.public.draft_points (draft_id, user_id, points)
VALUES ({{.DraftID}}, {{.UserID}}, {{.Points}})
ON CONFLICT (draft_id, user_id)
    DO UPDATE SET points = {{.Points}};