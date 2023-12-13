/*
 * 0200-0000-table-attachment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A collection of attachments for any ticket.
 * The attachment itself will use the attachmentId (uuid)
 * to query a global URL for the image content.
 */
create table if not exists attachment
(
    id       uuid primary key default gen_random_uuid(),
    ticketId uuid not null,
    authorId uuid not null,
    author   permissions not null default 'delete',
    team     permissions not null default 'read',
    everyone permissions not null default 'read',
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id)
)
