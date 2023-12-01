/*
 * 0007-table-attachment.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A collection of attachments for any ticket.
 */
create table if not exists attachment
(
    id       uuid primary key default gen_random_uuid(),
    ticketId uuid,
    authorId uuid,
    url      text,
    foreign key (ticketId) references ticket (id),
    foreign key (authorId) references users (id)
)
