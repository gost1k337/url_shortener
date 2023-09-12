CREATE TABLE IF NOT EXISTS users
(
    id            serial      not null primary key,
    username      varchar     not null unique,
    email         varchar     not null unique,
    password_hash varchar     not null,
    created_at    timestamptz not null default now()
)