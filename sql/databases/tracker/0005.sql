create table if not exists ticket
(
    id             uuid primary key default gen_random_uuid(),
    projectId      uuid          not null,
    authorUserId   uuid          not null,
    assignedUserId uuid          not null,
    subject        varchar(2048) not null,
    description    text
);
