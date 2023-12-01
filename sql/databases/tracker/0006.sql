/*
 * ticketAudience -
 *      Map tickets to users with permissions.
 */
create table if not exists ticketAudience
(
    id            uuid primary key default gen_random_uuid(),
    ticketId      uuid    not null,
    userId        uuid    not null,
    editorTicket  boolean not null default false, --edit the ticket
    deleteTicket  boolean not null default false, --delete the ticket
    createComment boolean not null default false,--create a comment
    viewComment   boolean not null default false, --view ticket comments
    editComment   boolean not null default false, --update ticket comments
    deleteComment boolean not null default false --delete ticket comments
)
