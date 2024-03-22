-- Returns all the friends for the user with the id == .UserID
SELECT DISTINCT u.id           AS id,
                u.display_name AS name
FROM public.friends f
         JOIN public.users u ON (CASE
                              WHEN f.user_id = {{.UserID}} THEN f.friend_id
                              ELSE f.user_id
    END) = u.id
WHERE (f.user_id = {{.UserID}} OR f.friend_id = {{.UserID}})
  AND f.relationship = 'friends';