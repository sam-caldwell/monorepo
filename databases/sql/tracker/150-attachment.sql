/*
 * 200-table-attachment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A collection of attachments for any ticket.
 * The attachment itself will use the attachmentId (uuid)
 * to query a global URL for the image content.
 */
create table if not exists attachment
(
    id       uuid primary key not null,
    ticketId uuid not null,
    authorId uuid not null,
    author   permissions not null default 'delete',
    team     permissions not null default 'read',
    everyone permissions not null default 'read',
    -- --
    created  timestamp        not null default now(),
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id),
    foreign key (id) references entity (id) on delete restrict
);

create index if not exists ndxAttachmentAuthor on attachment (author);
create index if not exists ndxAttachmentTeam on attachment (team);
create index if not exists ndxAttachmentEveryone on attachment (everyone);
