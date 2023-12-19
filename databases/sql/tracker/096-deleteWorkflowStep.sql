/*
 * 096-deleteWorkflowStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Perform any pre-checks that would preclude a record removal.
 * Update the prev and next steps to reference one another and remove the current node from the chain.
 * Delete the given workflow step from the database
 */
create or replace function deleteWorkflowStep(currentStepId uuid) returns integer as
$$
declare
    count  integer;
    prevId uuid;
    nextId uuid;
begin
    -- Get the current node's previous and next nodes...
    prevId := (select getWorkflowPrevStepId(currentStepId));
    nextId := (select getWorkflowNextStepId(currentStepId));
    -- Update previous and next node's next and prev references to exclude current node.
    execute linkNodes(prevId, nextId);
    -- Delete the current node
    delete from workflowSteps where id = currentStepId;
    get diagnostics count = ROW_COUNT;
    return count;
end ;
$$ language plpgsql;
