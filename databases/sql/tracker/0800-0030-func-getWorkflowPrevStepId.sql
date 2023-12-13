/*
 * 0800-0030-func-getWorkflowPrevStepId.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 */
create or replace function getWorkflowPrevStepId(stepId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', stepId,
            'name', name,
            'workflowId', workflowId,
            'iconId', iconId,
            'prevStepId', prevStepId,
            'nextStepId', nextStepId,
            'description', description
        )) as workflow
    into result
    from workflowSteps
    where id in (select prevStepId
                 from workflowSteps
                 where id = stepId);
    return result;
end;
$$ language plpgsql;
