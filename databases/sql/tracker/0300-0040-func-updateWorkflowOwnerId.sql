/*
 * 0008-0040-func-updateWorkflowOwnerId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowOwnerId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflow set ownerId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
