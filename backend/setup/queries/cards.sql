-- Table: public.cards

CREATE TABLE IF NOT EXISTS public.cards
(
    id integer NOT NULL,
    name text COLLATE pg_catalog."default",
    type text COLLATE pg_catalog."default",
    "desc" text COLLATE pg_catalog."default",
    atk integer,
    def integer,
    level integer,
    race text COLLATE pg_catalog."default",
    attribute text COLLATE pg_catalog."default",
    sets text COLLATE pg_catalog."default",
    CONSTRAINT cards_pkey PRIMARY KEY (id)
)