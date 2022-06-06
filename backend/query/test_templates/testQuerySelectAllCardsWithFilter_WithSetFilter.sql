SELECT "id", "name", "type", "desc", "atk", "def", "level", "race", "attribute", "sets" FROM public.cards
WHERE ("sets" like('%Force of the Breaker%') or "sets" like('%Strike of Neos%'))