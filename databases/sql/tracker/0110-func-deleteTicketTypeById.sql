/*
 * 0110-func-deleteTicketTypeById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteTicketTypeById(typeId uuid) returns integer as
$$
declare
    count integer;
begin
    -- We should not be able to delete a ticket type if it is in use.
    -- This means we cannot delete a ticket type if it is used in any workflow or ticket.
    delete from ticketTypes where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
