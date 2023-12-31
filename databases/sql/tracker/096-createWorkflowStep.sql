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
    /*
     * If our previous step (pStepId) is null, get the start node.
     */
    if (pStepId is null) then
        p := (select getStartNode(thisWorkflowId));
    else
        p := pStepId;
    end if;
    /*
     * If our next step (nStepId) is null, get the terminal node.
     */
    if (nStepId = '00000000-0000-0000-0000-000000000000'::uuid) then
        p := getTerminalNode(thisWorkflowId);
    else
        n := nStepId;
    end if;
    /*
     * Create our entityId.
     */
    stepId := (select createEntity('workflowStep'::entityType));
    if (stepId = pStepId) or (stepId = nStepId) then
        raise exception 'node cannot be its prev or next step id';
    end if;
    /*
     * insert the new workflow step record.
     */
    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (stepId, thisWorkflowId, stepName, p, n, stepDescription);
    /*
     * update previous node's next node pointer
     */
    update workflowsteps set nextstepid=stepId where workflowid=thisWorkflowId and id=p;
    /*
     * update the next node's prev node pointer
     */
    update workflowsteps set prevstepid=stepId where workflowid=thisWorkflowId and id=n;
    /*
     *  Now our records should create a double-linked list like this--
     *
     *      (p)<--->(stepId)<--->(n)
     */
    return stepId;
end;
$$ language plpgsql;
