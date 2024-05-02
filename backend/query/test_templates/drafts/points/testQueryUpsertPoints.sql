INSERT INTO ygodraft.public.draft_points (draft_id, user_id, points)
VALUES (4, 5, 55)
ON CONFLICT (draft_id, user_id)
    DO UPDATE SET points = 55;