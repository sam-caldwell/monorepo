/*
 * 0300-0040-func-updateWorkflowTeamId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowTeamId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflow set teamId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
