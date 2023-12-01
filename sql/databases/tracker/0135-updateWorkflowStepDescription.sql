/*
 * 0135-func-updateWorkflowStepDescription.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateWorkflowStepDescription(stepId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set description=description where id=stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
