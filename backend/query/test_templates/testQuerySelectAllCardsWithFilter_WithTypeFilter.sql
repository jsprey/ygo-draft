SELECT "id", "name", "type", "desc", "atk", "def", "level", "race", "attribute", "sets" FROM public.cards
WHERE "type" in('Normal Monster','Effect Monster')