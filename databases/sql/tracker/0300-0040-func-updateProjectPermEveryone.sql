/*
 * 0115-func-updateProjectPermEveryone
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectPermEveryone(projectId uuid, everyone permissions) returns integer as
$$
declare
    count integer;
begin
    update projects set everyone=everyone where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
