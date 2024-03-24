SELECT f.relationship
FROM public.friends as f
WHERE (f.user_id = {{.FromUserID}} AND f.friend_id = {{.ToUserID}}) OR (f.user_id = {{.ToUserID}} AND f.friend_id = {{.FromUserID}})
LIMIT 1