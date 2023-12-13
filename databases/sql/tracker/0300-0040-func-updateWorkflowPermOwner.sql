/*
 * 0300-0040-func-updateWorkflowPermOwner.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowPermOwner(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflow set owner=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
