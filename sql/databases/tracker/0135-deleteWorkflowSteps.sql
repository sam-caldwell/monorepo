/*
 * 0135-func-deleteWorkflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function deleteWorkflowSteps(stepId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from workflowSteps where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
