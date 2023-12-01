/*
 * 0100-func-deleteIcons.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function deleteIcons(iconId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from icons where id = iconId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
