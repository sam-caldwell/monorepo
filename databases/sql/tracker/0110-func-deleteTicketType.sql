/*
 * 0110-func-deleteTicketType.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteTicketType(typeId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from ticketTypes where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
