BEGIN;

INSERT INTO public.app_user (
	uid,
	username, 
	email,
    biography,
	password,
	created_at, 
	updated_at
)
VALUES (
    '1kF1ZTVJ3xknUgLWAAmPRKX3r8X',
    'testaccount_1',
    'testaccount_1@test.com',
    'biography',
    '$2a$04$MO1v6z7S8FNOj3GjY7yIzOMw6sfgMCTdLknXz4jMuFBGJA6CmI9zC',
    '2020-11-13 15:33:00.972651',
    '2020-11-13 15:33:00.972651'
);

INSERT INTO public.app_user (
	uid,
	username, 
	email,
	password,
	created_at, 
	updated_at
)
VALUES (
    '1kF488tsFU3qQkuXKWwUrPEI79K',
    'testaccount_2',
    'testaccount_2@test.com',
    'biography',
    '$2a$04$"1kF488tsFU3qQkuXKWwUrPEI79K"',
    '2020-11-13 15:54:03.662978',
    '2020-11-13 15:54:03.662978'
);

COMMIT;