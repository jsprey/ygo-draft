SELECT "id", "display_name", "email", "password_hash", "is_admin"
FROM public.users
ORDER BY "id"
OFFSET 20 ROWS
FETCH NEXT 20 ROWS ONLY;