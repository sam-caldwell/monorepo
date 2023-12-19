/*
 * 091-workflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 */
create table if not exists workflowSteps
(
    id          uuid primary key not null,
    workflowId  uuid             not null,
    name        varchar(64)      not null,
    -- a uuid representing the icon for the workflow.
    prevStepId  uuid             null,
    nextStepId  uuid             null,
    -- --
    created     timestamp        not null default now(),
    description text,
    foreign key (workflowId) references workflows (id),
    /*
     * see 093-verifyWorkFlowPrevNextStepsValid.sql
     *
     * We have two triggers (below) which enforce referential integrity
     * to ensure prevStepId and nextStepId are in the set of workflowSteps.id
     * if non-null uuid, but taking no action if the prevStepId or nextStepId
     * are null.
     */
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowStepsWorkflowIdName on workflowSteps (workflowId, name);
create index if not exists ndxWorkflowStepsCreated on workflowSteps (created);

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowStepId() function
 *      Get the current full step record as JSON object.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowPrevStepId() function
 *      Get the previous step Id
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowNextStepId() function
 *      Get the next step Id
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowNextStepId(stepId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select nextStepId
    into result
    from workflowSteps
    where id = stepId
    limit 1;
    return coalesce(result, '{}'::jsonb);
end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * linkNodes() function
 *      * Given two nodes (LHS and RHS) representing a prev node (left hand) and next node (right hand)
 *        this function will point lhs.nextId->rhs and rhs.prevId->lhs.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function linkNodes(lhs uuid, rhs uuid) returns boolean as
$$
begin
    execute updateWorkflowNextStep(lhs, rhs);
    execute updateWorkflowPrevStep(rhs, lhs);
    return true;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowStep() function
 *      Perform any pre-checks that would preclude a record removal.
 *      Update the prev and next steps to reference one another and remove the current node from the chain.
 *      Delete the given workflow step from the database
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateStepDescription(stepId uuid, d text) returns integer as
$$
declare
    count integer;
begin
    update teams
    set description=d
    where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowNextStep() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowNextStep(stepId uuid, nextStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set nextStepId=nextStepId where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowPrevStep() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowStepName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowStepName(stepId uuid, newName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set name=newName where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
