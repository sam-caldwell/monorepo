create table if not exists attachment
(
    id       uuid primary key default gen_random_uuid(),
    ticketId uuid,
    author   uuid,
    url      text
)
