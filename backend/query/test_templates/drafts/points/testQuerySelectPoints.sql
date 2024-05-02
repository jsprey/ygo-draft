SELECT d.id, d.draft_id, d.user_id, d.points
FROM ygodraft.public.draft_points AS d
WHERE draft_id = 4
  and user_id = 3