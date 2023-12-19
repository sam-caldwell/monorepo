/*
 * 096-updateWorkflowStepName.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create or replace function updateWorkflowStepName(stepId uuid, newName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set name=newName where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
