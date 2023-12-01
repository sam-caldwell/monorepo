create table if not exists comment (
    id uuid primary key default gen_random_uuid(),
    ticketId uuid,
    author uuid,
    comment text
)
