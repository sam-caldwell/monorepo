/*
 * 0135-func-updateWorkflowNextStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateWorkflowNextStep(stepId uuid, nextStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set nextStepId=nextStepId where id=stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
