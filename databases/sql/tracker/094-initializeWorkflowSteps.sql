/*
 * 094-initializeWorkflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create the start and terminate nodes of the workflow.
 */
create or replace function initializeWorkflowSteps(thisWorkflow uuid) returns integer as
$$
declare
    startId uuid := (select createEntity('workflow_step'::entityType));
    stopId  uuid := (select createEntity('workflow_step'::entityType));
begin
    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (startId, thisWorkflow, 'start', stopId, null, 'workflow starts here');

    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (stopId, thisWorkflow, 'terminate', null, startId, 'workflow stops here');

    return 1;
end ;
$$ language plpgsql;
