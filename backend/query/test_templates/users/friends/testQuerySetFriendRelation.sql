-- Update the status of the users relation
INSERT INTO public.friends (user_id, friend_id, relationship)
VALUES (5, 6, 'friends')
ON CONFLICT (user_id, friend_id)
    DO UPDATE SET relationship = 'friends';