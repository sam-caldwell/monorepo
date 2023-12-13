/*
 * 0800-0040-func-updateWorkflowPrevStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */

create or replace function updateWorkflowPrevStep(stepId uuid, prevStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set nextStepId=prevStepId where id=stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
