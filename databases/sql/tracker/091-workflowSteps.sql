/*
 * 091-workflowSteps.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create user profiles for the tracker.
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
 * check_foreign_key_constraints()
 *      Trigger on insert/update to ensure prevStepId and nextStepId are in workflowSteps
 *      unless they are null.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function check_foreign_key_constraints()
    returns trigger as
$$
declare
    prevNotInSet boolean;
    nextNotInSet boolean;
begin
    prevNotInSet := (select count(id) from workflowSteps where new.prevStepId in (select id from workflowSteps)) = 0;
    nextNotInSet := (select count(id) from workflowSteps where new.nextStepId in (select id from workflowSteps)) = 0;
    if (new.prevStepId is not null) and prevNotInSet then
        RAISE EXCEPTION 'Foreign key constraint violation: prevStepId must reference existing workflowSteps.id';
    end if;
    IF (new.nextStepId is not null) and nextNotInSet then
        RAISE EXCEPTION 'Foreign key constraint violation: nextStepId must reference existing workflowSteps.id';
    end if;
    return new;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_insert
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger check_foreign_key_insert
    before insert
    on workflowSteps
    for each row
execute function check_foreign_key_constraints();
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * check_foreign_key_update
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create trigger check_foreign_key_update
    before update
    on workflowSteps
    for each row
execute function check_foreign_key_constraints();

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * initializeWorkflowSteps() function
 *      Create the start and terminate nodes of the workflow.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function initializeWorkflowSteps(thisWorkflow uuid) returns jsonb as
$$
declare
    startId uuid := (select createEntity('workflow_step'::entityType));
    stopId  uuid := (select createEntity('workflow_step'::entityType));
begin
    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (startId, thisWorkflow, 'start', null, null, 'workflow starts here');

    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (stopId, thisWorkflow, 'terminate', null, null, 'workflow stops here');

    return jsonb_build_object(
            'startId', startId,
            'stopId', stopId);
end ;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createWorkflowStep() function
 *      Create a workflow step. Make sure we validate the data.
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createWorkflowStep(stepName varchar(64), thisworkflow uuid,
                                              pStepId uuid default null, nStepId uuid default null,
                                              stepDescription text default '') returns uuid as
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
        p := (select id from workflowsteps where workflowid = thisworkflow and name = 'start');
    else
        p := pStepId;
    end if;
    if (nStepId = '00000000-0000-0000-0000-000000000000'::uuid) then
        p := (select id from workflowsteps where workflowid = thisworkflow and name = 'start');
    else
        n := nStepId;
    end if;
    insert into workflowSteps (id, workflowId, name, prevStepId, nextStepId, description)
    values (stepId, workflow, stepName, p, n, stepDescription);
    return stepId;
end;
$$ language plpgsql;
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
