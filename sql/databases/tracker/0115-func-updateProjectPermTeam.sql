/*
 * 0115-func-updateProjectPermTeam
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectPermTeam(projectId uuid, team permissions) returns integer as
$$
declare
    count integer;
begin
    update projects set team=team where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
