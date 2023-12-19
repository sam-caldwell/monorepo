/*
 * 096-createWorkflowStep.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a workflow step and insert it between the previous step and next step.
 * if a null value is passed in, we will use the default start/terminal nodes.
 *
 * To create a workflow step...
 *      1. We assume there exists (start) <---> (terminate) at minimum.
 *      2. Thus if we pass null prevStepId or null nextStepId, we can assume a start
 *         or terminate.
 *      3. But if non-null values are provided, we should insert our step between the
 *         start or terminate.
 *      4. We should validate the stepName.
 */
create or replace function createWorkflowStep(stepName varchar(64), thisWorkflowId uuid,
                                              pStepId uuid, nStepId uuid,
                                              stepDescription text) returns uuid as
$$
declare
    stepId uuid;
    p      uuid;
    n      uuid;
begin
    stepId := (select createEntity('workflow_step'::entityType));
    if (stepId = pStepId) or (stepId = nStepId) then
        raise exception 'node cannot be its prev or next step id';
    end if;
    if (pStepId is null) then
        p := (select getStartNode(thisWorkflowId));
    else
        p := pStepId;
    end if;
    if (nStepId = '00000000-0000-0000-0000-000000000000'::uuid) then
        p := getTerminalNode(thisWorkflowId);
    else
        n := nStepId;
    end if;
    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (stepId, thisWorkflowId, stepName, p, n, stepDescription);
    return stepId;
end;
$$ language plpgsql;
