SELECT f.relationship
FROM public.friends as f
WHERE (f.user_id = 4 AND f.friend_id = 6) OR (f.user_id = 6 AND f.friend_id = 4)
LIMIT 1