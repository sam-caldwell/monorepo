/*
 * 130-workflowSteps.sql
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
    prevStepId  uuid             not null,
    nextStepId  uuid             not null,
    -- --
    created     timestamp        not null default now(),
    description text,
    foreign key (workflowId) references workflows (id),
    foreign key (prevStepId) references workflowSteps (id),
    foreign key (nextStepId) references workflowSteps (id),
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
 * createWorkflowSteps() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createWorkflowSteps(stepName varchar(64), workflow uuid, pStepId uuid, nStepId uuid,
                                               stepDescription text) returns uuid as
$$
declare
    teamId uuid;
begin
    teamId := (select createEntity('workflow_step'::entityType));
    insert into users (id, workflowId, name, prevStepId, nextStepId, description)
    values (teamId, workflow, stepName, pStepId, nStepId, stepDescription);
    return teamId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowSteps() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowSteps(stepId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from workflowSteps where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectDescription(stepId uuid) returns jsonb as
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
    where id == stepId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowPrevStepId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
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
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowNextStepId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowNextStepId(stepId uuid) returns jsonb as
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
    where id in (select nextStepId
                 from workflowSteps
                 where id = stepId);
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateTicketWorkflowStep() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateTicketWorkflowStep(ticketId uuid, workflowStepId uuid) returns integer as
$$
declare
    count integer;
begin
    update ticket set workflowStepId=workflowStepId where id = ticketId;
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
 * updateWorkflowStepDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowStepDescription(stepId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set description=description where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowStepDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowStepDescription(stepId uuid, iconId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set iconId=iconId where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowStepName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */

create or replace function updateWorkflowStepName(stepId uuid, name varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflowSteps set name=name where id = stepId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
