/*
 * 090-workflows.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This is the top-level table for representing a workflows.
 * This workflows is then composed of records in workflowsteps.
 */
create table if not exists workflows
(

    id          uuid primary key not null, -- workflowsId uniquely identifies the workflows --
    name        varchar(64)      not null, -- each workflows has a unique name --
    -- --
    iconId      uuid             not null, -- a uuid representing the icon for the workflows. --
    ownerId     uuid             not null, -- a workflows will have an owner and team which have permissions --
    teamId      uuid             not null, -- permissions are granted to the owner, team and everyone --
    -- --
    owner       permissions      not null default 'delete',
    team        permissions      not null default 'read',
    everyone    permissions      not null default 'read',
    -- --
    created     timestamp        not null default now(),
    -- --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (iconId) references icons (id),
    foreign key (id) references entity (id) on delete restrict
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxWorkflowsName on workflows (name);
create index if not exists ndxWorkflowsCreated on workflows (created);
create index if not exists ndxWorkflowsOwner on workflows (owner);
create index if not exists ndxWorkflowsTeam on workflows (team);
create index if not exists ndxWorkflowsEveryone on workflows (everyone);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createWorkflow()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createWorkflow(workflowName varchar(64), workflowIconId uuid, workflowOwnerId uuid,
                                          workflowTeamId uuid, workflowPermissionOwner permissions,
                                          workflowPermissionTeam permissions, workflowPermissionEveryone permissions,
                                          workflowPermissionDescription text) returns uuid as
$$
declare
    workflowsId uuid;
begin
    workflowsId := (select createEntity('workflow'::entityType));
    insert into workflows (id, name, iconId, ownerId, teamId, owner, team, everyone, description)
    values (workflowsId, workflowName, workflowIconId, workflowOwnerId, workflowTeamId,
            workflowPermissionOwner, workflowPermissionTeam, workflowPermissionEveryone,
            workflowPermissionDescription);
    return workflowsId;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowById(workflowId uuid) returns integer as
$$
declare
    count integer;
begin
    -- ToDo: we should not be able to delete any workflows if it is mapped to a ticketType
    delete from workflows where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowByName(workflowName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    -- ToDo: we should not be able to delete any workflows if it is mapped to a ticketType
    delete from workflows where name = workflowName;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowById(workflowId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'name', name,
                   'iconId', iconId,
                   'ownerId', ownerId,
                   'teamId', teamId,
                   'owner', owner,
                   'team', team,
                   'everyone', everyone,
                   'description', description
               ) as workflows
    into result
    from workflows
    where id = workflowId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowByName(workflowName varchar(64)) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_build_object(
                   'id', id,
                   'name', name,
                   'iconId', iconId,
                   'ownerId', ownerId,
                   'teamId', teamId,
                   'owner', owner,
                   'team', team,
                   'everyone', everyone,
                   'description', description
               ) as workflows
    into result
    from workflows
    where name = workflowName;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsByOwnerId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsByOwnerId(workflowOwnerId uuid,
                                                 pageLimit integer,
                                                 pageOffset integer) returns jsonb as
$$
declare
    discard bool;
    result  jsonb;
begin
    discard := (select boundsCheck(pageLimit, 1, 1000));
    discard := (select boundsCheck(pageOffset, 0, 1000));
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', teamId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'description', description
        )) as workflows
    into result
    from workflows
    where ownerId = workflowOwnerId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsByTeamId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsByTeamId(workflowTeamId uuid,
                                                pageLimit integer,
                                                pageOffset integer) returns jsonb as
$$
declare
    discard bool;
    result  jsonb;
begin
    discard := (select boundsCheck(pageLimit, 1, 1000));
    discard := (select boundsCheck(pageOffset, 0, 1000));
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', teamId,
            'owner', owner,
            'team', team,
            'everyone', everyone,
            'description', description
        )) as workflows
    into result
    from workflows
    where teamId = workflowTeamId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsDescription()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsDescription(workflowId uuid, descriptionText text) returns integer as
$$
declare
    count integer;
begin
    update workflows set odescriptionwnerId=descriptionText where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsIconId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsIconId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set iconId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsName(workflowId uuid, workflowName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflows set name=workflowName where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsOwnerId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsOwnerId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set ownerId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermEveryone()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermEveryone(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set everyone=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermOwner()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermOwner(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set owner=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermTeam(workflowId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set team=newPermission where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsTeamId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsTeamId(workflowId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set teamId=newId where id = workflowId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
