/*
 * 0140-func-updateWorkflowPermEveryone.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowPermEveryone(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflow set everyone=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
