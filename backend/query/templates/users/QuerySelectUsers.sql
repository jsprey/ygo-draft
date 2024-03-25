SELECT "id", "display_name", "email", "password_hash", "is_admin"
FROM public.users
ORDER BY "id"
OFFSET {{.Offset}} ROWS
FETCH NEXT {{.PageSize}} ROWS ONLY;