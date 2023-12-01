/*
 * 0105-func-updateTicketTypeName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketTypeName(typeId uuid, name varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set name=name where id=typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
