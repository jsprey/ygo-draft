-- Table: public.card_sets

CREATE TABLE IF NOT EXISTS public.card_sets
(
    set_code text COLLATE pg_catalog."default",
    set_name text COLLATE pg_catalog."default",
    set_rarity text COLLATE pg_catalog."default",
    set_rarity_code text COLLATE pg_catalog."default",
    CONSTRAINT card_sets_pkey PRIMARY KEY (set_name)
)