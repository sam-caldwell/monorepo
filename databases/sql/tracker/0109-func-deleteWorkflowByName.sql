/*
 * 0109-func-deleteWorkflowByName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteWorkflowByName(workflowName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    delete from workflow where name=workflowName;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
