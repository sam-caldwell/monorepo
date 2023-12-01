/*
 * 0140-func-deleteWorkflow.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteWorkflow(workflowId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from workflow where id=workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
