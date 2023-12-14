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
    name        varchar(64) not null, -- each workflows has a unique name --
    -- --
    iconId      uuid        not null, -- a uuid representing the icon for the workflows. --
    ownerId     uuid        not null, -- a workflows will have an owner and team which have permissions --
    teamId      uuid        not null, -- permissions are granted to the owner, team and everyone --
    -- --
    owner       permissions not null default 'delete',
    team        permissions not null default 'read',
    everyone    permissions not null default 'read',
    -- --
    created     timestamp   not null default now(),
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
create unique index if not exists ndxworkflowsName on workflows (name);
create index if not exists ndxworkflowsCreated on workflows (created);
create index if not exists ndxworkflowsOwner on workflows (owner);
create index if not exists ndxworkflowsTeam on workflows (team);
create index if not exists ndxworkflowsEveryone on workflows (everyone);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createWorkflows()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createWorkflows(name varchar(64), iconId uuid, ownerId uuid, teamId uuid, owner permissions,
                                          team permissions, everyone permissions, description text) returns uuid as
$$
declare
    workflowsId uuid;
begin
    workflowsId := (select createEntity('workflows'::entityType));
    insert into workflows (id, name, iconId, ownerId, teamId, owner, team, everyone, description)
    values (workflowsId, name, iconId, ownerId, teamId, owner, team, everyone, description);
    return workflowsId;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowsById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowsById(workflowsId uuid) returns integer as
$$
declare
    count integer;
begin
    -- ToDo: we should not be able to delete any workflows if it is mapped to a ticketType
    delete from workflows where id=workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteWorkflowsByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteWorkflowsByName(workflowsName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    -- ToDo: we should not be able to delete any workflows if it is mapped to a ticketType
    delete from workflows where name=workflowsName;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsById()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsById(workflowsId uuid) returns jsonb as
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
    where id = workflowsId;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsByName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsByName(workflowsName varchar(64)) returns jsonb as
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
    where name = workflowsName;
    return result;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsByOwnerId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsByOwnerId(workflowsOwnerId uuid,
                                                 pageLimit integer,
                                                 pageOffset integer) returns jsonb as
$$
declare
    discard bool;
    result jsonb;
begin
    discard:=(select boundsCheck(pageLimit,1,1000));
    discard:=(select boundsCheck(pageOffset,0,1000));
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
    where ownerId = workflowsOwnerId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getWorkflowsByTeamId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getWorkflowsByTeamId(workflowsTeamId uuid,
                                                pageLimit  integer,
                                                pageOffset integer) returns jsonb as
$$
declare
    discard bool;
    result jsonb;
begin
    discard:=(select boundsCheck(pageLimit,1,1000));
    discard:=(select boundsCheck(pageOffset,0,1000));
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
    where teamId = workflowsTeamId
    limit pageLimit offset pageOffset;
    return result;
end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsDescription()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsDescription(workflowsId uuid, description text) returns integer as
$$
declare
    count integer;
begin
    update workflows set odescriptionwnerId=descriptionText where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsIconId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsIconId(workflowsId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set iconId=newId where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsName()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsName(workflowsId uuid, workflowsName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update workflows set name=workflowsName where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsOwnerId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsOwnerId(workflowsId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set ownerId=newId where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermEveryone()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermEveryone(workflowsId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set everyone=newPermission where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermOwner()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermOwner(workflowsId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set owner=newPermission where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsPermTeam()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsPermTeam(workflowsId uuid, newPermission permissions) returns integer as
$$
declare
    count integer;
begin
    update workflows set team=newPermission where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateWorkflowsTeamId()
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateWorkflowsTeamId(workflowsId uuid, newId uuid) returns integer as
$$
declare
    count integer;
begin
    update workflows set teamId=newId where id = workflowsId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
