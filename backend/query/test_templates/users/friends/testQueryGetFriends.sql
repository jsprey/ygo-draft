-- Returns all the friends for the user with the id == .UserID
SELECT DISTINCT u.id           AS id,
                u.display_name AS name
FROM public.friends f
         JOIN public.users u ON (CASE
                              WHEN f.user_id = 5 THEN f.friend_id
                              ELSE f.user_id
    END) = u.id
WHERE (f.user_id = 5 OR f.friend_id = 5)
  AND f.relationship = 'friends';