/*
 * 0115-func-updateProjectIconId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectIconId(projectId uuid, iconId uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set iconId=iconId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
