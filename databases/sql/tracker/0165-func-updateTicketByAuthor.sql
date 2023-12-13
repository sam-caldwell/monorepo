/*
 * 0145-func-updateTicketAuthor.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketAuthor(ticketId uuid, authorUserId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticket set authorUserId=authorUserId where id=ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
