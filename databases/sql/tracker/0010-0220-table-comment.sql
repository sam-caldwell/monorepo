/*
 * 0010-0770-table-comment.sql
 * (c) 2023 Sam Caldwell. See License.txt.
 *
 * A collection of comments for any ticket.
 */
create table if not exists comment
(
    id       uuid primary key not null,
    ticketId uuid not null,
    authorId uuid not null,
    author   permissions not null default 'delete',
    team     permissions not null default 'read',
    everyone permissions not null default 'read',
    -- --
    created  timestamp        not null default now(),
    comment  text not null,
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id),
    foreign key (id) references entity (id) on delete restrict
);

create index if not exists ndxCommentAuthor on comment (author);
create index if not exists ndxCommentTeam on comment (team);
create index if not exists ndxCommentEveryone on comment (everyone);
