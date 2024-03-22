-- Returns all friend requests for the user with the user id UserID
SELECT u.id AS inviter_id,
       u.display_name AS inviter_name,
       f.invite_date AS invite_date
FROM public.friends f
         JOIN public.users u ON f.user_id = u.id
WHERE f.friend_id = 5
  AND f.relationship = 'invited';