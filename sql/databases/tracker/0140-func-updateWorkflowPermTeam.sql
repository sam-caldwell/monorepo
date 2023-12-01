/*
 * 0140-func-updateWorkflowPermTeam.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowPermTeam(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflow set team=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
