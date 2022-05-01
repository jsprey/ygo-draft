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
    image_small bytea,
    image_big bytea,
    sets integer[],
    CONSTRAINT cards_pkey PRIMARY KEY (id)
)