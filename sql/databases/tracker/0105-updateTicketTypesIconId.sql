/*
 * 0105-func-updateTicketTypeIconId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateTicketTypeIconId(typeId uuid, iconId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticketTypes set iconid=iconId where id=typeId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

