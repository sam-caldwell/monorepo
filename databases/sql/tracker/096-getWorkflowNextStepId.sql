/*
 * 096-getWorkflowNextStepId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Get the next step Id
 */
create or replace function getWorkflowNextStepId(stepId uuid) returns uuid as
$$
declare
    result uuid;
begin
    select nextStepId
    into result
    from workflowSteps
    where id = stepId
    limit 1;
    return result;
end ;
$$ language plpgsql;
