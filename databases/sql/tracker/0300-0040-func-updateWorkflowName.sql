/*
 * 0008-0040-func-updateWorkflowName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowName(workflowId uuid, workflowName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflow set name=workflowName where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
