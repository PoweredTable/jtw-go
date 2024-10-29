create table public."user"
(
    id            integer generated always as identity primary key,
    name          varchar(150) not null,
    password_hash text         not null,
    phone         varchar(20),
    email         varchar(150) not null unique,
    role          user_role default 'user'::user_role,
    created_at    timestamp default CURRENT_TIMESTAMP,
    is_active     boolean   default true
);

alter table public."user"
    owner to postgres;