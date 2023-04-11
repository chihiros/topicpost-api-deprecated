CREATE TABLE profile (
	id int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY,
	nickname text NOT NULL,
	uuid uuid NOT NULL,
	icon_url text NULL,
	created_at timestamptz NULL DEFAULT now(),
	updated_at timestamptz NULL DEFAULT now(),
);

INSERT INTO profile
    (nickname, uuid, icon_url)
VALUES
    (
        'suzurikawa',
        'c2cc015d-5753-4132-86fe-624787ae49df',
        'https://avatars.githubusercontent.com/u/1000000001'
    ),(
        'tanaka',
        '87cab8ea-2252-4a4d-bc15-3fe6192bbf5d',
        'https://avatars.githubusercontent.com/u/1000000002'
    );
