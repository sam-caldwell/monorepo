/*
 * 0115-func-updateProjectTeamId
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateProjectTeamId(projectId uuid, teamId uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set team=teamId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
