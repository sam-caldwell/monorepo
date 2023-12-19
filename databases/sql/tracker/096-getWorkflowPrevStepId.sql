/*
 * 096-getWorkflowPrevStepId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Get the previous step Id
 */
create or replace function getWorkflowPrevStepId(stepId uuid) returns uuid as
$$
declare
    result jsonb;
begin
    select prevStepId
    into result
    from workflowSteps
    where id = stepId
    limit 1;
    return coalesce(result, '{}'::jsonb);
end ;
$$ language plpgsql;
