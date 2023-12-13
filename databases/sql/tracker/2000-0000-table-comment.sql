/*
 * 2000-0000-table-comment.sql
 * (c) 2023 Sam Caldwell. See License.txt.
 *
 * A collection of comments for any ticket.
 */
create table if not exists comment
(
    id       uuid primary key     default gen_random_uuid(),
    ticketId uuid not null,
    authorId uuid not null,
    author   permissions not null default 'delete',
    team     permissions not null default 'read',
    everyone permissions not null default 'read',
    comment  text not null,
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id)
);
