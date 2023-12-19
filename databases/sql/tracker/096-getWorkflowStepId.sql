/*
 * 091-getWorkflowStepId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Get the current full step record as JSON object.
 */
create or replace function getWorkflowStepId(stepId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'workflowId', workflowId,
                   'name', name,
                   'prevStepId', prevStepId,
                   'nextStepId', nextStepId,
                   'created', created,
                   'description', description
               ) as data
    into result
    from workflowSteps
    where id = stepId
    limit 1;
    return result::jsonb;
end ;
$$ language plpgsql;
