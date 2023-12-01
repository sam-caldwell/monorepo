create table if not exists teams
(
    id          uuid primary key default gen_random_uuid(),
    name        varchar(64)      not null unique,
    owner       uuid             not null,
    description text
);
