BEGIN;

CREATE TABLE public.user
(
    id SERIAL PRIMARY KEY,
    uid VARCHAR (36) NOT NULL UNIQUE,
    username VARCHAR (15) NOT NULL UNIQUE,
    email VARCHAR (254) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_on timestamp NOT NULL,
    updated_on timestamp NOT NULL
);

COMMIT;
