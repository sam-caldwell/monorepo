/*
 * 1000-0020-func-deleteTicket.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteTicket(ticketId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from ticket where id = ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
