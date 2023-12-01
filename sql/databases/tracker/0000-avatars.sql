create table if not exists avatars
(
    id  uuid primary key default gen_random_uuid(),
    url text
);
