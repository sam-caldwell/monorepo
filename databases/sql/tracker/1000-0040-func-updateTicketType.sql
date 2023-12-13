/*
 * 1000-0040-func-updateTicketType.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketType(ticketId uuid, ticketTypeId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticket set ticketTypeId=ticketTypeId where id=ticketId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
