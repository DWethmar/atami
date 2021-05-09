BEGIN;

CREATE TABLE public.app_user
(
    id SERIAL PRIMARY KEY,
    uid VARCHAR (36) NOT NULL UNIQUE,
    username VARCHAR (15) NOT NULL UNIQUE,
    biography text,
    email VARCHAR (254) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE public.message
(
    id SERIAL PRIMARY KEY,
    uid VARCHAR (36) NOT NULL UNIQUE,
    text TEXT NOT NULL,
    created_by_user_id integer REFERENCES public.app_user (id) ON DELETE CASCADE,
    created_at timestamp NOT NULL
    updated_at timestamp NOT NULL
);

COMMIT;
