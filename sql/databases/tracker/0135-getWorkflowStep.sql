/*
 * 0135-func-getWorkflowStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 */
create or replace function getWorkflowStep(stepId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', stepId,
            'name', name,
            'workflowId', workflowId,
            'iconId',iconId,
            'prevStepId',prevStepId,
            'nextStepId',nextStepId,
            'description',description
        )) as workflow
    into result
    from workflowSteps
    where id == stepId;
    return result;
end;
$$ language plpgsql;
