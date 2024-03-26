DELETE
FROM public.friends as f
WHERE f.friend_id = {{.UserID}} or f.user_id = {{.UserID}};