create table if not exists statuses
(
    id          uuid primary key default gen_random_uuid(),
    name        varchar(32) not null unique,
    description text
);
