/*
 * 110-projects.sql
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 */
create table if not exists projects
(
    id          uuid primary key not null,
    -- each workflow has a unique name --
    name        varchar(64)      not null,
    -- a uuid representing the icon for the workflow. --
    iconId      uuid             not null,
    -- a project will have an owner and team which have permissions --
    ownerId     uuid             not null,
    teamId      uuid             not null,
    -- permissions --
    owner       permissions      not null default 'delete',
    team        permissions      not null default 'none',
    everyone    permissions      not null default 'none',
    -- --
    created     timestamp        not null default now(),
    -- descriptive text --
    description text,
    foreign key (ownerId) references users (id),
    foreign key (teamId) references teams (id),
    foreign key (iconId) references icons (id),
    foreign key (id) references entity (id) on delete restrict,
    constraint validateProjectName check (validName(name))
);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * entity indexes
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create unique index if not exists ndxProjectsName on projects (name);
create index if not exists ndxProjectsCreated on projects (created);
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * createProject() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function createProject(projectName varchar(64), projectIconId uuid, projectOwnerId uuid,
                                         projectTeamId uuid, permissionOwner permissions, permissionTeam permissions,
                                         permissionEveryone permissions, projectDesc text) returns uuid as
$$
declare
    newId uuid;
begin
    newId := (select createEntity('project'::entityType));
    insert into projects (id, name, iconId, ownerId, teamId, owner, team, everyone, description)
    values (newId, projectName, projectIconId, projectOwnerId, projectTeamId,
            permissionOwner, permissionTeam, permissionEveryone, projectDesc);
    return newId;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * deleteProject() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function deleteProject(projectId uuid) returns integer as
$$
declare
    count integer;
begin
    delete from projects where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;

/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getProjectById() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getProjectById(projectId uuid) returns jsonb as
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
                   'permissionOwner', owner,
                   'permissionTeam', team,
                   'permissionEveryone', everyone,
                   'created', created,
                   'description', description
               ) as data
    into result
    from projects
    where id == projectId
    limit 1;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getProjectByName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getProjectByName(projectName varchar(64)) returns jsonb as
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
                   'permissionOwner', owner,
                   'permissionTeam', team,
                   'permissionEveryone', everyone,
                   'created', created,
                   'description', description
               ) as data
    into result
    from projects
    where name = projectName
    limit 1;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getProjectsByOwnerId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getProjectsByOwnerId(thisOwnerId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', thisOwnerId,
            'teamId', teamId,
            'permissionOwner', owner,
            'permissionTeam', team,
            'permissionEveryone', everyone,
            'created', created,
            'description', description
        )) as data
    into result
    from projects
    where ownerId = thisOwnerId;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * getProjectsByTeamId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function getProjectsByTeamId(thisTeamId uuid) returns jsonb as
$$
declare
    result jsonb;
begin
    select jsonb_agg(jsonb_build_object(
            'id', id,
            'name', name,
            'iconId', iconId,
            'ownerId', ownerId,
            'teamId', thisTeamId,
            'permissionOwner', owner,
            'permissionTeam', team,
            'permissionEveryone', everyone,
            'created', created,
            'description', description
        )) as data
    into result
    from projects
    where teamId = thisTeamId;
    return result;

end ;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectDescription(projectId uuid, descriptiveText text) returns integer as
$$
declare
    count integer;
begin
    update projects set description=descriptiveText where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectIconId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectIconId(projectId uuid, projectIconId uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set iconId=projectIconId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectName() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectName(projectId uuid, projectName varchar(64)) returns integer as
$$
declare
    count integer;
begin
    update projects set name=projectName where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectOwnerId() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectOwnerId(projectId uuid, newOwnerId uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set ownerId=newOwnerId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectPermissions() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectPermissions(projectId uuid, pOwner permissions, pTeam permissions,
                                                    pEveryone permissions) returns integer as
$$
declare
    count integer;
begin
    update projects set owner=pOwner, team=pTeam, everyone=pEveryone where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
/*
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * updateProjectDescription() function
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 */
create or replace function updateProjectTeamId(projectId uuid, newTeamId uuid) returns integer as
$$
declare
    count integer;
begin
    update projects set teamId=newTeamId where id = projectId;
    get diagnostics count = ROW_COUNT;
    return count;
end;
$$ language plpgsql;
