/*
 * 0300-0020-func-deleteWorkflowById.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteWorkflowById(workflowId uuid) returns integer as
$$
declare
    count integer;
begin
    -- we should not be able to delete any workflow if it is mapped to a project
    delete from workflow where id=workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
