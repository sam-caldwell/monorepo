create table if not exists workflow
(
    id          uuid primary key     default gen_random_uuid(),
    name        varchar(64) not null unique,
    enabled     bool        not null default true,
    description text
)
