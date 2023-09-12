CREATE TABLE IF NOT EXISTS url_shorts (
    id serial not null primary key,
    original_url varchar not null,
    short_url varchar not null,
    visits int not null default 0,
    expire_at timestamptz not null,
    created_at timestamptz not null default now()
)