/*
 * 152-deleteTicket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteTicket(ticketId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from tickets where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
