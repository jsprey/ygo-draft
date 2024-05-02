SELECT d.id, d.draft_id, d.user_id, d.points
FROM ygodraft.public.draft_points AS d
WHERE draft_id = {{.DraftID}}
  and user_id = {{.UserID}}