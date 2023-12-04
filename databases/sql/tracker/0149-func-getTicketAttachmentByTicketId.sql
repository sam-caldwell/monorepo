/*
 * 0148-func-getTicketAttachmentByTicket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
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
        )) as workflow into result
    from attachment
    where ticketId == ticketId
    limit pageLimit offset pageOffset;
    return result;
end;
$$ language plpgsql;
