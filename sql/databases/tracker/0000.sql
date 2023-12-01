create table if not exists users
(
    id           uuid primary key default gen_random_uuid(),
    teamId       uuid             not null,
    first_name   varchar(64)      not null,
    last_name    varchar(64)      not null,
    email        varchar(128)     not null unique,
    phone_number varchar(20)      not null,
    description  text
);
