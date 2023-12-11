/*
 * 0125-func-updateProjectOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectOwnerId(projectId uuid, ownerId varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update projects set ownerId=ownerId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
