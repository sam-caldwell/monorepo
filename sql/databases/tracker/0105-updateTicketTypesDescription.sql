/*
 * 0105-func-updateTicketTypeDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketTypeDescription(typeId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set description=description where id = typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
