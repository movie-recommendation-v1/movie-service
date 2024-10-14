create table if not exists movies(
    id uuid primary key not null,
    moviename varchar(350) not null,
    agelimit int not null,
    season int not null,
    backround_image_url varchar(350) not null,
    movie_url varchar(350) not null,
    studio varchar(250) not null,
    bio text not null,
    genres text[] not null,
    languege varchar(55) not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at bigint default 0 not null
);

create table if not exists comments(
    id uuid primary key not null,
    user_id uuid not null,
    movie_id uuid references movies(id),
    description text not null,
    rate int not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at bigint default 0 not null
);