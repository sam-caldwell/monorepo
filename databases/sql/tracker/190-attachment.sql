/*
 * 190-attachment.sql
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
    foreign key (ticketId) references tickets (id),
    foreign key (authorId) references users (id),
    foreign key (id) references entity (id) on delete restrict
);

create index if not exists ndxAttachmentAuthor on attachment (author);
create index if not exists ndxAttachmentTeam on attachment (team);
create index if not exists ndxAttachmentEveryone on attachment (everyone);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createTicketAttachment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createTicketAttachment(ticketId uuid, authorId uuid,
                                                  permAuthor permissions, permTeam permissions,
                                                  permEveryone permissions) returns uuid as
$$
declare
    commentId uuid;
begin
    commentId := (select createEntity('attachment'::entityType));
    insert into attachment (id, ticketId, authorId, permAuthor, permTeam, permEveryone, everyone)
    values (commentId, ticketId, authorId, permAuthor, permTeam, permEveryone, url);
    return commentId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteTicketAttachment()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteTicketAttachment(commentId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from attachment where id = commentId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getTicketAttachmentByTicket()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getTicketAttachmentByTicket(ticketId uuid,
                                                       pageLimit integer,
                                                       pageOffset integer) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'ticketId', ticketId,
            'authorId', authorId,
            'author', author,
            'team', team,
            'everyone', everyone
        )) as workflow
    into result
    from attachment
    where ticketId == ticketId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
