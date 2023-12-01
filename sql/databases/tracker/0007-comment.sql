create table if not exists comment
(
    id       uuid primary key     default gen_random_uuid(),
    ticketId uuid,
    authorId uuid,
    author   permissions not null default 'delete',
    everyone permissions not null default 'read',
    comment  text,
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id)
);
