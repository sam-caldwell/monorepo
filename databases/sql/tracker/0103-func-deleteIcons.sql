/*
 * 0103-func-deleteIcons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteIcons(iconId uuid) returns integer as
$$
declare
    count integer;
begin
    -- an icon cannot be deleted if it is in use (project, workflow, ticketType, team, etc)
    delete from icons where id = iconId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
