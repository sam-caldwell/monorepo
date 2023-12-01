create table if not exists icons
(
    id  uuid primary key default gen_random_uuid(),
    url text
);
