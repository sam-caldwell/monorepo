/*
 * 096-updateWorkflowPrevStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Perform any pre-checks that would preclude a record removal.
 * Update the prev and next steps to reference one another and remove the current node from the chain.
 * Delete the given workflow step from the database
 */
create or replace function updateWorkflowPrevStep(stepId uuid, prevStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set nextStepId=prevStepId where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
