SELECT "id", "display_name", "email", "password_hash", "is_admin" FROM public.users WHERE "id" = {{.ID}} LIMIT 1