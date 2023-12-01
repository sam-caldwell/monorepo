create table if not exists projects
(
    id          uuid primary key default gen_random_uuid(),
    name        varchar(64) not null unique,
    owner       uuid        not null,
    team        uuid        not null,
    workflow    uuid        not null,
    description text
);
